package main

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := map[string]struct{
		input *TreeNode
		want int
	}{
		"Test Case 1": {
			input: &TreeNode{
				Val: 1,
				Left: nil,
				Right: &TreeNode{
					Val: 2,
					Left: nil,
					Right: nil,
				},
			},
			want: 2,
		},
		"Test Case 2": {
			input: &TreeNode{
				Val: 0,
				Left: nil,
				Right: nil,
			},
			want: 1},
		"Test Case 3": {input: nil, want: 0},
		"Test Case 4": {
			input: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
					Left: nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
						Left: nil,
						Right: nil,
					},
					Right: &TreeNode{
						Val: 7,
						Left: nil,
						Right: nil,
					},
				},
			},
			want: 3},
	}

	for name, tc := range tests {
		got := MaxDepth(tc.input)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}