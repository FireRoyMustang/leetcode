package main

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

// 链表中节点的数目在范围 [0, 500] 内
// -100 <= Node.val <= 100
// 0 <= k <= 2 * 10^9

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	length := 0
	node := head
	for node != nil {
		node = node.Next
		length++
	}
	// for node != nil && length < k {
	// 	node = node.Next
	// 	length++
	// }
	k = k % length
	// if length < k {
	// 	k = k % length
	// } else {

	// }
	if k == 0 {
		return head
	}
	node = head
	m := length - k
	for m > 1 {
		node = node.Next
		m--
	}
	newHead := node.Next
	node.Next = nil
	node = newHead
	for node.Next != nil {
		node = node.Next
	}
	node.Next = head
	return newHead
}
