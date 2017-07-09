package main

import (
	"testing"
)

func TestHashing(t *testing.T) {
	b := &Block{Content: "foo", PrevHash: [32]byte{}, Nonce: 42}
	s := [32]byte{136, 166, 247, 160, 108, 15, 71, 148, 105, 169, 67, 82, 60, 178, 228, 151, 180, 11, 129, 166, 96, 158, 99, 22, 122, 160, 119, 176, 118, 89, 233, 26}
	if b.Hash() != s {
		t.Error("Block did not generate the right hash")
	}
}
