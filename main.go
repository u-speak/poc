package main

import (
	"crypto/sha256"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	"github.com/u-speak/poc/util"
	"google.golang.org/grpc"
	"strconv"
	"time"
)

const version = "0.0.1"

func main() {
	err := configor.Load(&Config, "config.yml")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Starting uspeak poc version %s", version)
	log.SetLevel(log.DebugLevel)
	c := chain.New(validateBeginningZero)
	mineAndAdd("Bootstrap Block 1", c)
	mineAndAdd("Bootstrap Block 2", c)
	mineAndAdd("foo", c)
	mineAndAdd("bar", c)
	ns := &NodeServer{chain: c, remoteConnections: make(map[string]*grpc.ClientConn)}
	go ns.Run()
	// Connect the node to itself
	err = ns.Connect(Config.NodeNetwork.Interface + ":" + strconv.Itoa(Config.NodeNetwork.Port))
	if err != nil {
		log.Fatal(err)
	}
	repl(c, ns)
	log.Info("Shutting down by interactive command")
	ns.Shutdown()
}

func mineAndAdd(content string, c *chain.Chain) uint {
	nonce := mine(content, c.LastHash())
	logIfError(c.AddData(content, nonce))
	return nonce
}

func validateBeginningZero(h [32]byte) bool {
	return h[0] == 0
}

func logIfError(err error) {
	if err != nil {
		log.Error(err)
	}
}

func mine(content string, prev [32]byte) uint {
	start := time.Now()
	log.WithField("prev", util.CompactEmoji(prev)).Debugf("Started mining %s", content)
	var i uint
	for i = 0; i < ^uint(0); i++ {
		h := sha256.Sum256([]byte(string(prev[:]) + content + string(i)))
		if validateBeginningZero(h) {
			elapsed := time.Since(start)
			log.WithField("elapsed", elapsed).Debugf("Found Nonce %d for %s", i, content)
			return i
		}
	}
	panic("Block impossible to mine")
}
