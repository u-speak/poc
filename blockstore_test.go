package main

import (
	"testing"
)

func BenchmarkMemoryStore(b *testing.B) {
	s := &MemoryStore{}
	for i := 0; i < b.N; i++ {
		bl := &Block{Content: "foo"}
		s.Add(bl)
	}
}
