package linked_list_helper

type ListNode struct {
	Val int
	Next *ListNode
}

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