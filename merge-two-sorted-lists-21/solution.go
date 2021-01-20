package main

type ListNode struct {
    Val int
    Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    temp := ListNode{}

    if l1.Val < l2.Val {
        temp.Val = l1.Val
    }

    return &temp
}