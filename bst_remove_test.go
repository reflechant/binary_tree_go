package bst

import (
	"reflect"
	"testing"
)

func TestRemoveNonExistent(t *testing.T) {
	n1 := &Node{1, &Node{-2, nil, nil}, &Node{2, nil, &Node{3, nil, nil}}}
	n2 := &Node{1, &Node{-2, nil, nil}, &Node{2, nil, &Node{3, nil, nil}}}
	n1 = n1.Remove(7)
	if !Equals(n1, n2) {
		t.FailNow()
	}
}

func TestRemoveRoot(t *testing.T) {
	n1 := &Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}}
	n2 := &Node{3, &Node{1, nil, nil}, nil}
	n1 = n1.Remove(2)
	if !Equals(n1, n2) {
		t.FailNow()
	}
}

func TestRemoveGeneral(t *testing.T) {
	n1 := New(5, 3, 7, 2, 8, 1)
	n1 = n1.Remove(5)
	if !reflect.DeepEqual(n1.Keys(), []int{1, 2, 3, 7, 8}) {
		t.FailNow()
	}
}
