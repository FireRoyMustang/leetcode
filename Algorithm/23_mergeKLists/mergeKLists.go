package main

import "fmt"

//   Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	lists := make([]*ListNode, 0)
	// arr1, arr2, arr3 := []int{1, 4, 5}, []int{1, 3, 4}, []int{2, 6}
	// list1 := generateList(arr1)
	// list2 := generateList(arr2)
	// list3 := generateList(arr3)
	// lists = append(lists, list1)
	// lists = append(lists, list2)
	// lists = append(lists, list3)

	lists = append(lists, nil)
	lists = append(lists, nil)
	printList(mergeKLists(lists))

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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	helpNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	head := helpNode
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			helpNode.Next = l2
			l2 = l2.Next
		} else {
			helpNode.Next = l1
			l1 = l1.Next
		}
		helpNode = helpNode.Next
	}
	if l1 == nil {
		helpNode.Next = l2
	}
	if l2 == nil {
		helpNode.Next = l1
	}
	return head.Next
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) >> 1
	return mergeTwoLists(merge(lists, left, mid), merge(lists, mid+1, right))
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func mergeKLists1(lists []*ListNode) *ListNode {
	k := len(lists)
	if k == 0 {
		return nil
	}
	if k == 1 {
		return lists[0]
	}
	helpNode := &ListNode{
		Val:  -1,
		Next: nil,
	}
	isNilList := make([]bool, k)
	isNilNum := 0
	minNodeIndex, minNodeNum := 0, 0
	head := helpNode
	for isNilNum < k-1 {
		minNodeNum = 1 << 31
		for i := 0; i < k; i++ {
			if !isNilList[i] {
				if lists[i] == nil {
					isNilList[i] = true
					isNilNum++
				} else {
					if lists[i].Val < minNodeNum {
						minNodeNum = lists[i].Val
						minNodeIndex = i
					}
				}
			}
		}
		if isNilNum < k-1 {
			helpNode.Next = lists[minNodeIndex]
			helpNode = helpNode.Next
			lists[minNodeIndex] = lists[minNodeIndex].Next
		}

	}
	for index, flag := range isNilList {
		if !flag {
			helpNode.Next = lists[index]
		}
	}
	return head.Next
}
