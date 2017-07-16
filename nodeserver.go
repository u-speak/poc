package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	d "github.com/u-speak/poc/distribution"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
)

// NodeServer implements the DistributionServiceServer interface
type NodeServer struct {
	chain             *chain.Chain
	remoteConnections map[string]*grpc.ClientConn
}

// GetInfo implements the coresponding RPC call
func (s *NodeServer) GetInfo(ctx context.Context, params *d.StatusParams) (*d.Info, error) {
	if _, contained := s.remoteConnections[params.Host]; !contained {
		err := s.Connect(params.Host)
		if err != nil {
			log.Error("Failed to initialize reverse connection. Fulfilling request anyways...")
		}
	}
	lh := s.chain.LastHash()
	return &d.Info{
		Version:  version,
		Valid:    s.chain.IsValid(),
		Length:   s.chain.Length(),
		LastHash: lh[:],
	}, nil
}

// Synchronize streams the missing block to the partner node
func (s *NodeServer) Synchronize(p *d.SyncParams, stream d.DistributionService_SynchronizeServer) error {
	h := s.chain.LastHash()
	b := s.chain.Get(h)
	var c [32]byte
	copy(c[:], p.LastHash)
	last := s.chain.LastHash()
	for {
		if err := stream.Send(&d.Block{Content: b.Content, Nonce: uint32(b.Nonce), Previous: b.PrevHash[:], Last: last[:]}); err != nil {
			log.Error(err)
		}
		if b.PrevHash == c {
			break
		}
		b = s.chain.Get(b.PrevHash)
	}
	return nil
}

// Run starts the server
func (s *NodeServer) Run() {
	lis, _ := net.Listen("tcp", Config.NodeNetwork.Interface+":"+strconv.Itoa(Config.NodeNetwork.Port))
	grpcServer := grpc.NewServer()
	d.RegisterDistributionServiceServer(grpcServer, s)
	log.Fatal(grpcServer.Serve(lis))
}

// Connect adds a remote node to the list of known nodes
func (s *NodeServer) Connect(remote string) error {
	if _, contained := s.remoteConnections[remote]; contained {
		return errors.New("Node allready connected")
	}
	conn, err := grpc.Dial(remote, grpc.WithInsecure())
	if err != nil {
		return err
	}
	s.remoteConnections[remote] = conn
	log.Infof("Successfully connected to %s", remote)
	return nil
}

// SynchronizeChain adds missing blocks to the chain
func (s *NodeServer) SynchronizeChain(remote string) error {
	lh := s.chain.LastHash()
	params := &d.SyncParams{LastHash: lh[:]}
	client := d.NewDistributionServiceClient(s.remoteConnections[remote])
	stream, err := client.Synchronize(context.Background(), params)
	if err != nil {
		return err
	}
	for {
		block, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var p [32]byte
		copy(p[:], block.Previous)
		err = s.chain.AddRaw(&chain.Block{Content: block.Content, Nonce: uint(block.Nonce), PrevHash: p}, block.Last...)
		if err != nil {
			return err
		}
	}
	return nil
}

// Push pushes a block to all connected nodes
func (s *NodeServer) Push(b *chain.Block) {
	lh := s.chain.LastHash()
	for _, r := range s.remoteConnections {
		client := d.NewDistributionServiceClient(r)
		_, err := client.Receive(context.Background(), &d.Block{Content: b.Content, Nonce: uint32(b.Nonce), Previous: lh[:]})
		if err != nil {
			log.Error(err)
		}
	}
}

// Receive adds the pushed block to the chain
func (s *NodeServer) Receive(ctx context.Context, block *d.Block) (*d.PushReturn, error) {
	log.Debugf("Received Block: %s", block.Content)
	var p [32]byte
	copy(p[:], block.Previous)
	if p != s.chain.LastHash() {
		log.Errorf("Tried to add invalid Block! Previous hash %v is not valid. Please synchronize the nodes", p)
		return &d.PushReturn{}, errors.New("Received block had invalid previous hash")
	}
	return &d.PushReturn{}, s.chain.AddData(block.Content, uint(block.Nonce))
}

// Shutdown closes all remote Connections
func (s *NodeServer) Shutdown() {
	for _, r := range s.remoteConnections {
		_ = r.Close()
	}
}
