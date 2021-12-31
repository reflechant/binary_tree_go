package bst

import "testing"

func TestEquals(t *testing.T) {
	cases := []struct {
		name     string
		a        Node
		b        Node
		expected bool
	}{
		{"single nodes equality test",
			Node{1, nil, nil},
			Node{1, nil, nil},
			true},
		{"tree equality test",
			Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}},
			Node{2, &Node{1, nil, nil}, &Node{3, nil, nil}},
			true},
		{"nodes with different values are not equal",
			Node{1, nil, nil},
			Node{2, nil, nil},
			false,
		},
		{"single node not equals node with left child",
			Node{1, &Node{0, nil, nil}, nil},
			Node{1, nil, nil},
			false,
		},
		{"single node not equals node with right child",
			Node{1, nil, &Node{2, nil, nil}},
			Node{1, nil, nil},
			false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if Equals(&tc.a, &tc.b) != tc.expected ||
				Equals(&tc.b, &tc.a) != tc.expected {
				t.Fail()
			}
		})
	}
}
