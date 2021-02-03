package main

import (
	"fmt"
	"leetcode/tools"
)

// 从上到下打印出二叉树的每个节点，同一层的节点按照从左到右的顺序打印。
//   Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	tree := tools.GenerateTree(test)
	fmt.Println("Tree:", tools.LevelOrder(tree))
}

func levelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			node := queue[0]
			res = append(res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			size--
		}
	}
	return res
}
