package bst

import "testing"

func TestRemoveNonExistent(t *testing.T) {
	n1 := Node{1, &Node{-2, nil, nil}, &Node{2, nil, &Node{3, nil, nil}}}
	n2 := Node{1, &Node{-2, nil, nil}, &Node{2, nil, &Node{3, nil, nil}}}
	n1.Remove(7)
	if !Equals(&n1, &n2) {
		t.FailNow()
	}
}
