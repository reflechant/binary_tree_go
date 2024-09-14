package bst

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

func TestFindRoot(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, nil}
	n2 := n1.Find(1)
	assert.Equal(t, &n1, n2)
}

func TestFindLeftChild(t *testing.T) {
	n1 := Node[int, string]{1, "", &Node[int, string]{-1, "", nil, nil}, nil}
	n2 := n1.Find(-1)
	assert.Equal(t, n2, n1.Left)
}

func TestFindRightChild(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, &Node[int, string]{2, "", nil, nil}}
	n2 := n1.Find(2)
	assert.Equal(t, n2, n1.Right)
}

func TestFindNonExistent(t *testing.T) {
	n1 := Node[int, string]{1, "", nil, &Node[int, string]{2, "", nil, nil}}
	n2 := n1.Find(5)
	assert.Equal(t, nil, n2)
}
