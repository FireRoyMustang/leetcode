package main

// 定义一个函数，输入一个链表的头节点，反转该链表并输出反转后链表的头节点。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reverseList(head *ListNode) *ListNode {
	node := head
	var prevNode *ListNode = nil
	for node != nil {
		tmp := node.Next
		node.Next = prevNode
		prevNode = node
		node = tmp
	}
	return prevNode
}

// func reverseList(head *ListNode) *ListNode {
// 	if head == nil {
// 		return nil
// 	}
// 	node := head
// 	var prevNode *ListNode = nil
// 	for node.Next != nil {
// 		tmp := node.Next
// 		node.Next = prevNode
// 		prevNode = node
// 		node = tmp
// 	}
// 	node.Next = prevNode
// 	return node
// }
