package bst

import "testing"

func TestOneNodeMinMax(t *testing.T) {
	n1 := Node[int,string]{1, "", nil, nil}
	min := n1.Min()
	if min.Key != 1 {
		t.FailNow()
	}
	max := n1.Max()
	if max.Key != 1 {
		t.FailNow()
	}
}

func TestTwoLeavesMinMax(t *testing.T) {
	n1 := Node[int,string]{2, "", &Node[int,string]{1, "", nil, nil}, &Node[int,string]{3, "", nil, nil}}
	min := n1.Min()
	if min.Key != 1 {
		t.FailNow()
	}
	max := n1.Max()
	if max.Key != 3 {
		t.FailNow()
	}
}
