package bst

import (
	"testing"
)

func BenchmarkKeys(b *testing.B) {
	n := &Node[int, string]{5000, "", nil, nil}
	for i := range 10000 {
		n.Insert(i, "")
	}
	b.ResetTimer()
	var sum int
	for range b.N {
		for k := range n.Keys() {
			sum += k
		}
	}
}
