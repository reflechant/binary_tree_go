package bst

import "testing"

func TestAppend(t *testing.T) {
	cases := []struct {
		name     string
		a        Node
		b        int
		expected Node
	}{
		{"append right to single node",
			Node{1, nil, nil}, 2,
			Node{1, nil, &Node{2, nil, nil}}},
		{"append left to single node",
			Node{1, nil, nil}, 0,
			Node{1, &Node{0, nil, nil}, nil}},
		{"append duplicate",
			Node{1, nil, nil}, 1,
			Node{1, nil, nil}},
		{"append left to tree",
			Node{1, &Node{0, nil, nil}, nil}, -1,
			Node{1, &Node{0, &Node{-1, nil, nil}, nil}, nil}},
		{"append right to tree",
			Node{1, nil, &Node{2, nil, nil}}, 3,
			Node{1, nil, &Node{2, nil, &Node{3, nil, nil}}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.a.Append(tc.b)
			if !Equals(&tc.a, &tc.expected) {
				t.Fail()
			}
		})
	}
}
