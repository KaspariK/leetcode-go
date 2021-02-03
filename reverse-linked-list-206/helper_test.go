package reverse_linked_list_206

import (
	"reflect"
	"testing"
)

// SliceToLinkedList converts an []int to a ListNode. This just makes tests a bit easier to write and format
func SliceToLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	prev := head

	for _, n := range nums {
		curr := &ListNode{
			Val: n,
		}

		if prev != nil {
			prev.Next = curr
		}

		prev = curr
	}

	return head.Next
}

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
