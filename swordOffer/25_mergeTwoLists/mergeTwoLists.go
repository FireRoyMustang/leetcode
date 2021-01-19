package main

// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	var node *ListNode = head
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			node.Next = l2
			l2 = l2.Next
		} else {
			node.Next = l1
			l1 = l1.Next
		}
		node = node.Next
	}
	if l1 == nil {
		node.Next = l2
	} else {
		node.Next = l1
	}
	return head.Next
}
