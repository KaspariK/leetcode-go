package reverse_linked_list_206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prevNode *ListNode
	curNode := head

	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = prevNode
		prevNode = curNode
		curNode = nextNode
	}
	return prevNode
}
