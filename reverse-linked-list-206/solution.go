package reverse_linked_list_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	prevNode := nil
	curNode := head

	// 1->2->3->4->5->NIL
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	return head
}
