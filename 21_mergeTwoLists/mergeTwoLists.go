package main

// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	helpNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	head := helpNode
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			helpNode.Next = l2
			l2 = l2.Next
		} else {
			helpNode.Next = l1
			l1 = l1.Next
		}
		helpNode = helpNode.Next
	}
	if l1 == nil {
		helpNode.Next = l2
	}
	if l2 == nil {
		helpNode.Next = l1
	}
	return head.Next
}
