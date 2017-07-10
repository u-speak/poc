package tests

import (
	"github.com/u-speak/poc/chain"
	"strconv"
	"testing"
)

func BenchmarkValidationCheck(b *testing.B) {
	c := chain.New(func(h [32]byte) bool { return true })
	for i := 0; i < 10000; i++ {
		_ = c.AddData(strconv.Itoa(i), 0)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.IsValid()
	}
}

func BenchmarkPop(b *testing.B) {
	c := chain.New(func(h [32]byte) bool { return true })
	for i := 0; i < 10000; i++ {
		_ = c.AddData(strconv.Itoa(i), 0)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		c.Get(c.LastHash())
	}
}
