package bst

import "testing"

func slicesEqual(seq1, seq2 []int) bool {
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

func TestSingleNodeKeys(t *testing.T) {
	n := Node{Key: 1, Left: nil, Right: nil}
	if !slicesEqual(n.Keys(), []int{1}) {
		t.FailNow()
	}
}

func TestKeysOrder(t *testing.T) {
	n1 := Node{1, &Node{2, nil, nil}, nil}
	n2 := Node{1, nil, &Node{2, nil, nil}}
	if slicesEqual(n1.Keys(), n2.Keys()) {
		t.FailNow()
	}
}
