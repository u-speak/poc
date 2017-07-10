package chain

import (
	"crypto/sha256"
)

// Block is a concrete data entity
type Block struct {
	Content  string
	Nonce    uint
	PrevHash [32]byte
}

// Hash is the computed hash of the block
func (b *Block) Hash() [32]byte {
	return sha256.Sum256([]byte(string(b.PrevHash[:]) + b.Content + string(b.Nonce)))
}
