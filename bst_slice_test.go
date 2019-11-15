package bst

import "testing"

func eq(seq1, seq2 []int) bool {
	if len(seq1) != len(seq2) {
		return false
	}
	for i := 0; i < len(seq1); i++ {
		if seq1[i] != seq2[i] {
			return false
		}
	}
	return true
}

func TestSingleNodeSlice(t *testing.T) {
	n := Node{Key: 1, Left: nil, Right: nil}
	if !eq(n.Slice(), []int{1}) {
		t.FailNow()
	}
}

func TestSliceOrder(t *testing.T) {
	n1 := Node{1, &Node{2, nil, nil}, nil}
	n2 := Node{1, nil, &Node{2, nil, nil}}
	if eq(n1.Slice(), n2.Slice()) {
		t.FailNow()
	}
}
