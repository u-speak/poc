// Code generated by protoc-gen-go. DO NOT EDIT.
// source: distribution.proto

/*
Package distribution is a generated protocol buffer package.

It is generated from these files:
	distribution.proto

It has these top-level messages:
	Info
	StatusParams
	SyncParams
	Block
	PushReturn
*/
package distribution

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Info struct {
	Version  string `protobuf:"bytes,1,opt,name=Version" json:"Version,omitempty"`
	Valid    bool   `protobuf:"varint,2,opt,name=Valid" json:"Valid,omitempty"`
	Length   uint64 `protobuf:"varint,3,opt,name=Length" json:"Length,omitempty"`
	LastHash []byte `protobuf:"bytes,4,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
}

func (m *Info) Reset()                    { *m = Info{} }
func (m *Info) String() string            { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()               {}
func (*Info) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Info) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Info) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Info) GetLength() uint64 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Info) GetLastHash() []byte {
	if m != nil {
		return m.LastHash
	}
	return nil
}

type StatusParams struct {
	Host string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
}

func (m *StatusParams) Reset()                    { *m = StatusParams{} }
func (m *StatusParams) String() string            { return proto.CompactTextString(m) }
func (*StatusParams) ProtoMessage()               {}
func (*StatusParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StatusParams) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type SyncParams struct {
	LastHash []byte `protobuf:"bytes,1,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
}

func (m *SyncParams) Reset()                    { *m = SyncParams{} }
func (m *SyncParams) String() string            { return proto.CompactTextString(m) }
func (*SyncParams) ProtoMessage()               {}
func (*SyncParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SyncParams) GetLastHash() []byte {
	if m != nil {
		return m.LastHash
	}
	return nil
}

type Block struct {
	Content  string `protobuf:"bytes,1,opt,name=Content" json:"Content,omitempty"`
	Nonce    uint32 `protobuf:"varint,2,opt,name=Nonce" json:"Nonce,omitempty"`
	Previous []byte `protobuf:"bytes,3,opt,name=Previous,proto3" json:"Previous,omitempty"`
	Last     []byte `protobuf:"bytes,4,opt,name=Last,proto3" json:"Last,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Block) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Block) GetNonce() uint32 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Block) GetPrevious() []byte {
	if m != nil {
		return m.Previous
	}
	return nil
}

func (m *Block) GetLast() []byte {
	if m != nil {
		return m.Last
	}
	return nil
}

type PushReturn struct {
}

func (m *PushReturn) Reset()                    { *m = PushReturn{} }
func (m *PushReturn) String() string            { return proto.CompactTextString(m) }
func (*PushReturn) ProtoMessage()               {}
func (*PushReturn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func init() {
	proto.RegisterType((*Info)(nil), "Info")
	proto.RegisterType((*StatusParams)(nil), "StatusParams")
	proto.RegisterType((*SyncParams)(nil), "SyncParams")
	proto.RegisterType((*Block)(nil), "Block")
	proto.RegisterType((*PushReturn)(nil), "PushReturn")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DistributionService service

type DistributionServiceClient interface {
	GetInfo(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*Info, error)
	Synchronize(ctx context.Context, in *SyncParams, opts ...grpc.CallOption) (DistributionService_SynchronizeClient, error)
	Receive(ctx context.Context, in *Block, opts ...grpc.CallOption) (*PushReturn, error)
}

type distributionServiceClient struct {
	cc *grpc.ClientConn
}

func NewDistributionServiceClient(cc *grpc.ClientConn) DistributionServiceClient {
	return &distributionServiceClient{cc}
}

func (c *distributionServiceClient) GetInfo(ctx context.Context, in *StatusParams, opts ...grpc.CallOption) (*Info, error) {
	out := new(Info)
	err := grpc.Invoke(ctx, "/DistributionService/GetInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *distributionServiceClient) Synchronize(ctx context.Context, in *SyncParams, opts ...grpc.CallOption) (DistributionService_SynchronizeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DistributionService_serviceDesc.Streams[0], c.cc, "/DistributionService/Synchronize", opts...)
	if err != nil {
		return nil, err
	}
	x := &distributionServiceSynchronizeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DistributionService_SynchronizeClient interface {
	Recv() (*Block, error)
	grpc.ClientStream
}

type distributionServiceSynchronizeClient struct {
	grpc.ClientStream
}

func (x *distributionServiceSynchronizeClient) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *distributionServiceClient) Receive(ctx context.Context, in *Block, opts ...grpc.CallOption) (*PushReturn, error) {
	out := new(PushReturn)
	err := grpc.Invoke(ctx, "/DistributionService/Receive", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DistributionService service

type DistributionServiceServer interface {
	GetInfo(context.Context, *StatusParams) (*Info, error)
	Synchronize(*SyncParams, DistributionService_SynchronizeServer) error
	Receive(context.Context, *Block) (*PushReturn, error)
}

func RegisterDistributionServiceServer(s *grpc.Server, srv DistributionServiceServer) {
	s.RegisterService(&_DistributionService_serviceDesc, srv)
}

func _DistributionService_GetInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributionServiceServer).GetInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributionService/GetInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributionServiceServer).GetInfo(ctx, req.(*StatusParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _DistributionService_Synchronize_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SyncParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DistributionServiceServer).Synchronize(m, &distributionServiceSynchronizeServer{stream})
}

type DistributionService_SynchronizeServer interface {
	Send(*Block) error
	grpc.ServerStream
}

type distributionServiceSynchronizeServer struct {
	grpc.ServerStream
}

func (x *distributionServiceSynchronizeServer) Send(m *Block) error {
	return x.ServerStream.SendMsg(m)
}

func _DistributionService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DistributionServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DistributionService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DistributionServiceServer).Receive(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

var _DistributionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "DistributionService",
	HandlerType: (*DistributionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetInfo",
			Handler:    _DistributionService_GetInfo_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _DistributionService_Receive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Synchronize",
			Handler:       _DistributionService_Synchronize_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "distribution.proto",
}

func init() { proto.RegisterFile("distribution.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xbb, 0xef, 0x9b, 0xfe, 0x71, 0x9a, 0x5e, 0x46, 0x91, 0x90, 0x53, 0xdc, 0x83, 0xe4,
	0x14, 0x44, 0xbf, 0x81, 0x0a, 0x56, 0x28, 0x52, 0xb6, 0xd0, 0x7b, 0x9a, 0x8e, 0xcd, 0x6a, 0xdd,
	0x95, 0xdd, 0x4d, 0x41, 0x8f, 0x7e, 0x72, 0xd9, 0x6d, 0xda, 0x06, 0x6f, 0xfb, 0x1b, 0xc2, 0x33,
	0xbf, 0x67, 0x02, 0xb8, 0x96, 0xd6, 0x19, 0xb9, 0x6a, 0x9c, 0xd4, 0xaa, 0xf8, 0x34, 0xda, 0x69,
	0xfe, 0x06, 0xd1, 0xb3, 0x7a, 0xd5, 0x98, 0xc0, 0x70, 0x49, 0xc6, 0x4a, 0xad, 0x12, 0x96, 0xb1,
	0xfc, 0x4c, 0x1c, 0x10, 0x2f, 0xa0, 0xbf, 0x2c, 0xb7, 0x72, 0x9d, 0xfc, 0xcb, 0x58, 0x3e, 0x12,
	0x7b, 0xc0, 0x4b, 0x18, 0xcc, 0x48, 0x6d, 0x5c, 0x9d, 0xfc, 0xcf, 0x58, 0x1e, 0x89, 0x96, 0x30,
	0x85, 0xd1, 0xac, 0xb4, 0x6e, 0x5a, 0xda, 0x3a, 0x89, 0x32, 0x96, 0xc7, 0xe2, 0xc8, 0x9c, 0x43,
	0xbc, 0x70, 0xa5, 0x6b, 0xec, 0xbc, 0x34, 0xe5, 0x87, 0x45, 0x84, 0x68, 0xaa, 0xad, 0x6b, 0x17,
	0x86, 0x37, 0xcf, 0x01, 0x16, 0x5f, 0xaa, 0x6a, 0xbf, 0xe8, 0xa6, 0xb1, 0x3f, 0x69, 0x1b, 0xe8,
	0xdf, 0x6f, 0x75, 0xf5, 0xee, 0xd5, 0x1f, 0xb4, 0x72, 0xa4, 0x0e, 0x49, 0x07, 0xf4, 0xea, 0x2f,
	0x5a, 0x55, 0x14, 0xd4, 0x27, 0x62, 0x0f, 0x3e, 0x74, 0x6e, 0x68, 0x27, 0x75, 0x63, 0x83, 0x7c,
	0x2c, 0x8e, 0xec, 0x95, 0xfc, 0x82, 0x56, 0x3d, 0xbc, 0x79, 0x0c, 0x30, 0x6f, 0x6c, 0x2d, 0xc8,
	0x35, 0x46, 0xdd, 0xfe, 0x30, 0x38, 0x7f, 0xec, 0xdc, 0x71, 0x41, 0x66, 0x27, 0x2b, 0xc2, 0x2b,
	0x18, 0x3e, 0x91, 0x0b, 0xb7, 0x9c, 0x14, 0xdd, 0x9a, 0x69, 0xbf, 0xf0, 0x53, 0xde, 0xc3, 0x6b,
	0x18, 0xfb, 0x6e, 0xb5, 0xd1, 0x4a, 0x7e, 0x13, 0x8e, 0x8b, 0x53, 0xd3, 0x74, 0x50, 0x84, 0x32,
	0xbc, 0x77, 0xc3, 0x30, 0x83, 0xa1, 0xa0, 0x8a, 0xe4, 0x8e, 0xb0, 0x1d, 0xa7, 0xe3, 0xe2, 0xa4,
	0xc0, 0x7b, 0xab, 0x41, 0xf8, 0x79, 0x77, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x77, 0xcb, 0xea,
	0x80, 0xd2, 0x01, 0x00, 0x00,
}
