package bst

import "testing"

func TestNewNodeEqualsLiteral(t *testing.T) {
	n1 := NewNode(0)
	n2 := Node{0, nil, nil}
	if !Equals(&n1, &n2) || !Equals(&n2, &n1) {
		t.FailNow()
	}
}
