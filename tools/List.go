package tools

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func generateList(arr []int) *ListNode {
	head := &ListNode{arr[0], nil}
	node := head
	for i := 1; i < len(arr); i++ {
		node.Next = &ListNode{arr[i], nil}
		node = node.Next
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
}
