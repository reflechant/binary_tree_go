package bst

import "testing"

func TestNewNodeEqualsLiteral(t *testing.T) {
	n1 := New(0)
	n2 := &Node{0, nil, nil}
	if !Equals(n1, n2) {
		t.FailNow()
	}
}
