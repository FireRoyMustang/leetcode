package main
import "fmt"
// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
// 返回同样按升序排列的结果链表。


func main(){
	printList(deleteDuplicates(generateList([]int{1,2,3,3,4,4,5})))
	printList(deleteDuplicates(generateList([]int{1,1,1,2,3})))
}
func printList(node *ListNode) {
	for node != nil {
		fmt.Printf("%d ", node.Val)
		node = node.Next
	}
	fmt.Println()
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
type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	helpNode := &ListNode{-1, head}
	pre := helpNode
	flag := false // 节点是否相等
	for head != nil{
		// 删除重复节点
		for head.Next != nil && head.Next.Val == head.Val {
			head = head.Next
			flag = true
		}
		if flag {
			pre.Next = head.Next
			flag = false
		}else{
			pre = head
		}
		head = head.Next
	}
	return helpNode.Next
}
/* func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	helpNode := &ListNode{-1, head}
	pre := helpNode
	for head != nil{
		// 删除重复节点
		if head.Next != nil && head.Next.Val == head.Val{ // 有重复节点
			for head.Next != nil && head.Next.Val == head.Val {
				head = head.Next
			}
			pre.Next = head.Next
			head = head.Next
		}else { // 无重复节点
			pre = head
			head = head.Next
		}
	}
	return helpNode.Next
} */