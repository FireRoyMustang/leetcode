package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	node := head
	for node.Next != nil {
		for node.Val == node.Next.Val {
			node.Next = node.Next.Next
			if node.Next == nil {
				return head
			}
		}
		node = node.Next
	}
	return head
}
