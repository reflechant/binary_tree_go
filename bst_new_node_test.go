package bst

import (
	"reflect"
	"testing"
)

func TestNewNodeEqualsLiteral(t *testing.T) {
	n1 := New(0)
	n2 := &Node{0, nil, nil}
	if !Equals(n1, n2) {
		t.FailNow()
	}
}

func TestNewNodeFromMultipleKeys(t *testing.T) {
	n1 := New(2, 1, 3)
	n2 := &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}}
	if !reflect.DeepEqual(n1, n2) {
		t.FailNow()
	}
}
