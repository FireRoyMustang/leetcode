package main

import "fmt"

// 给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。
// k 是一个正整数，它的值小于或等于链表的长度。
// 如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

func main() {
	test := generateList([]int{1, 2, 3, 4, 5})
	head := reverseKGroup(test, 3)
	printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode, k int) (headd, tail, next *ListNode) {
	hair := &ListNode{-1, head}
	var preNode *ListNode = nil
	count := 0
	for ; count < k && head != nil; count++ {
		// fmt.Println(count, head.Val)
		tmp := head.Next
		head.Next = preNode
		preNode = head
		head = tmp
	}
	if count < k {
		preNode, _, _ = reverseList(preNode, count)
	}
	return preNode, hair.Next, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{-1, head}
	node, preNode := head, hair
	for node != nil {
		tmpHead, tmpTail, next := reverseList(node, k)
		preNode.Next = tmpHead
		preNode = tmpTail
		node = next
	}
	return hair.Next
}
func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Print("\n")
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
