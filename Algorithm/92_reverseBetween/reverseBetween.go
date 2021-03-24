package main

import "fmt"

// 链表中节点数目为 n
// 1 <= n <= 500
// -500 <= Node.val <= 500
// 1 <= left <= right <= n

func main() {
	head := &ListNode{3, nil}
	head.Next = &ListNode{5, nil}
	head = reverseBetween(head, 1, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	helpNode := &ListNode{-1, head}
	var pre *ListNode = helpNode
	for i := 1; i < left; i++ {
		pre = pre.Next
	}
	startNode := pre
	tailNode := pre.Next
	cur := pre.Next.Next
	pre = pre.Next
	for i := left; i < right; i++ {
		tmp := cur.Next
		cur.Next = pre
		fmt.Println(cur.Next.Val)
		pre = cur
		cur = tmp
	}
	// fmt.Printf("pre: %d\n", pre.Val)
	// fmt.Printf("pre.Next: %d\n", pre.Next.Val)
	// fmt.Printf("pre: %d\n", cur.Val)
	startNode.Next = pre
	tailNode.Next = cur
	return helpNode.Next
}
