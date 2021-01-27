package reverse_linked_list_206

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := map[string]struct {
		input *ListNode
		want  *ListNode
	}{
		"Test Case 1": {
			input: &ListNode{
				Val: 1,
				Next: nil,
			},
			want: &ListNode{
				Val: 1,
				Next: nil,
			},
		},
		"Test Case 2": {
			input: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: nil,
					},
				},
			},
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
		got := reverseList(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
