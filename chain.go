package main

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/theSuess/uspeak-poc/util"
)

// ValidationFunc is the requirement for mining
type ValidationFunc func([32]byte) bool

// Chain is a Blockchain Implementation
type Chain struct {
	blocks   BlockStore
	lastHash [32]byte
	validate ValidationFunc
}

// New initializes a new Chain
func New(validate ValidationFunc) *Chain {
	return &Chain{blocks: &MemoryStore{}, validate: validate}
}

// AddData adds a new block to the chain
func (c *Chain) AddData(content string, nonce uint) error {
	b := &Block{Content: content, PrevHash: c.lastHash, Nonce: nonce}
	hash := b.Hash()
	if !c.validate(hash) {
		return errors.New("Block did not pass the validation function")
	}
	c.blocks.Add(b)
	c.lastHash = hash
	return nil
}

// PrintChain logs the chain for debugging purposes
func (c *Chain) PrintChain() error {
	if !c.IsValid() {
		return errors.New("Chain is not Valid! Cannot print")
	}
	h := c.lastHash
	for h != [32]byte{} {
		b := c.printBlock(h)
		h = b.PrevHash
	}
	return nil
}

func (c *Chain) printBlock(hash [32]byte) *Block {
	b := c.blocks.Get(hash)
	log.WithFields(log.Fields{
		"hash":     util.CompactEmoji(b.Hash()),
		"prevHash": util.CompactEmoji(b.PrevHash),
	}).Debug(b.Content)
	return b
}

// Get retrieves a block
func (c *Chain) Get(hash [32]byte) *Block {
	return c.blocks.Get(hash)
}

// IsValid checks the chain for integrity and validation compliance
func (c *Chain) IsValid() bool {
	if c.lastHash == [32]byte{} {
		return true
	}
	b := c.blocks.Get(c.lastHash)
	for b != nil {
		if b.PrevHash == [32]byte{} {
			return true
		}
		if !c.validate(b.Hash()) {
			return false
		}
		b = c.blocks.Get(b.PrevHash)
	}
	return false
}

// LastHash returns the hash of the last block in the chain
func (c *Chain) LastHash() [32]byte {
	return c.lastHash
}
