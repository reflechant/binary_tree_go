package bst

import "testing"

func TestInsert(t *testing.T) {
	cases := []struct {
		name     string
		a        Node
		b        int
		expected Node
	}{
		{"insert right to single node",
			Node{1, nil, nil}, 2,
			Node{1, nil, &Node{2, nil, nil}}},
		{"insert left to single node",
			Node{1, nil, nil}, 0,
			Node{1, &Node{0, nil, nil}, nil}},
		{"insert duplicate",
			Node{1, nil, nil}, 1,
			Node{1, nil, nil}},
		{"insert left to tree",
			Node{1, &Node{0, nil, nil}, nil}, -1,
			Node{1, &Node{0, &Node{-1, nil, nil}, nil}, nil}},
		{"insert right to tree",
			Node{1, nil, &Node{2, nil, nil}}, 3,
			Node{1, nil, &Node{2, nil, &Node{3, nil, nil}}}},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			tc.a.Insert(tc.b)
			if !Equals(&tc.a, &tc.expected) {
				t.Fail()
			}
		})
	}
}
