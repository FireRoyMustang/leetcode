package main

// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func reversePrint(head *ListNode) []int {
	length := 0
	var node *ListNode = head
	for node != nil {
		length++
		node = node.Next
	}
	ans := make([]int, length)
	for head != nil {
		length--
		ans[length] = head.Val
		head = head.Next
	}
	return ans
}
