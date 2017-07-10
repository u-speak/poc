package main

import (
	"crypto/sha256"
	log "github.com/sirupsen/logrus"
	"github.com/u-speak/poc/chain"
	"github.com/u-speak/poc/util"
	"time"
)

const version = "0.0.1"

func main() {
	log.Infof("Starting uspeak poc version %s", version)
	log.SetLevel(log.DebugLevel)
	c := chain.New(validateBeginningZero)
	mineAndAdd("Bootstrap Block 1", c)
	mineAndAdd("Bootstrap Block 2", c)
	mineAndAdd("foo", c)
	mineAndAdd("bar", c)
	repl(c)
}

func mineAndAdd(content string, c *chain.Chain) {
	logIfError(c.AddData(content, mine(content, c.LastHash())))
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
