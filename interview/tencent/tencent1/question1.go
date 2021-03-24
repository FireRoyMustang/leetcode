package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, nil}
	// root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}},
	// 	&TreeNode{3, &TreeNode{6, nil, nil}, &TreeNode{7, nil, nil}}}
	nodes := solve(root, []int{1, 2, 3})
	for _, node := range nodes {
		for node != nil {
			fmt.Printf("%d\t", node.Val)
			node = node.Next
		}
		fmt.Println()
	}
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 你需要返回m个指针，第i个指针指向一条链，表示第i个问题的答案
 * @param root TreeNode类 指向链表树的根
 * @param b int整型一维数组 表示每个问题是什么
 * @return ListNode类一维数组
 */
func solve(root *TreeNode, b []int) []*ListNode {
	// write code here
	if root == nil {
		return []*ListNode{}
	}
	parents := make(map[int]*ListNode)
	res := make([]*ListNode, 0)
	queue := []*TreeNode{root}
	size := len(queue)
	var head *ListNode = nil
	for size != 0 {
		for i := 0; i < size; i++ { // 每个树节点
			node := queue[0]
			queue = queue[1:]
			parentNode := parents[node.Val]
			if parentNode != nil {
				head = &ListNode{parentNode.Val, nil}
				helpNode := head
				for parentNode.Next != nil {
					helpNode.Next = &ListNode{parentNode.Next.Val, nil}
					parentNode = parentNode.Next
					helpNode = helpNode.Next
				}
				helpNode.Next = &ListNode{node.Val, nil}
			} else {
				head = &ListNode{node.Val, nil}
			}
			if node.Left != nil {
				parents[node.Left.Val] = head
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				parents[node.Right.Val] = head
				queue = append(queue, node.Right)
			}
			res = append(res, head)
		}
		size = len(queue)
	}
	return res
}
