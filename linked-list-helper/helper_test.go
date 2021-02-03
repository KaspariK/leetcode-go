package linked_list_helper

import (
	"reflect"
	"testing"
)

func TestSliceToLinkedList(t *testing.T) {
	tests := map[string]struct{
		input []int
		want *ListNode
	}{
		"Test Case 1": {input: []int{1, 2, 3}, want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}},
		"Test Case 2": {input: nil, want: nil},
		"Test Case 3": {input: []int{1}, want: &ListNode{Val: 1, Next: nil}},
		"Test Case 4": {input: []int{1, 1}, want: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}},
	}

	for name, tc := range tests {
		got := SliceToLinkedList(tc.input)

		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("%s: expected '%v', got '%v'", name, tc.want, got)
		}
	}
}
