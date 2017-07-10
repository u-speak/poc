package tests

import (
	"github.com/u-speak/poc/chain"
	"testing"
)

func BenchmarkMemoryStore(b *testing.B) {
	s := &chain.MemoryStore{}
	for i := 0; i < b.N; i++ {
		bl := &chain.Block{Content: "foo"}
		s.Add(bl)
	}
}
