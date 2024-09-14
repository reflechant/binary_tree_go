package bst

import (
	"testing"
)

func BenchmarkKeys(b *testing.B) {
	n := &Node[int, string]{50000, "", nil, nil}
	for i := range 100000 {
		n.Insert(i, "")
	}
	b.ResetTimer()
	var sum int
	for range b.N {
		for k := range n.Keys() {
			sum += k
		}
	}
	q = sum
}

var q int

func BenchmarkKeys2(b *testing.B) {
	n := &Node[int, string]{50000, "", nil, nil}
	for i := range 100000 {
		n.Insert(i, "")
	}
	b.ResetTimer()
	var sum int
	for range b.N {
		for k := range n.Keys2() {
			sum += k
		}
	}
	q = sum
}
