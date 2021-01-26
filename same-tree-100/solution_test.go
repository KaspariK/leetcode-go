package same_tree_100

import (
	"testing"
)

func TestIsSameTree(t *testing.T) {
	tests := map[string]struct {
		treeOne TreeNode
		treeTwo TreeNode
		want    bool
	}{
		"Test Case 1": {
			TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   3,
					Left:  nil,
					Right: nil,
				},
			},
			true,
		},
		"Test Case 2": {
			TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			},
			TreeNode{
				Val:  1,
				Left: nil,
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			false,
		},
		"Test Case 3": {
			TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
			},
			TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val:   1,
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   2,
					Left:  nil,
					Right: nil,
				},
			},
			false,
		},
	}

	for name, tc := range tests {
		got := isSameTree(&tc.treeOne, &tc.treeTwo)

		if got != tc.want {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
