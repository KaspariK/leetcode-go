package reverse_linked_list_206

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  *ListNode
	}{
		"Test Case 1": {
			input: []int{1},
			want: &ListNode{
				Val: 1,
				Next: nil,
			},
		},
		"Test Case 2": {
			input: []int{1, 2, 3},
			want: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: nil,
					},
				},
			},
		},
		"Test Case 3": {input: nil, want: nil},
	}

	for name, tc := range tests {
		a := SliceToLinkedList(tc.input)
		got := reverseList(a)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
