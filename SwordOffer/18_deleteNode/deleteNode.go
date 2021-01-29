package main

// 给定单向链表的头指针和一个要删除的节点的值，定义一个函数删除该节点。
// 返回删除后的链表的头节点。

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	helpNode := &ListNode{
		Val:  -1,
		Next: head,
	}
	node, prevNode := head, helpNode
	for node.Val != val {
		prevNode = node
		node = node.Next
	}
	prevNode.Next = node.Next
	return helpNode.Next
}
