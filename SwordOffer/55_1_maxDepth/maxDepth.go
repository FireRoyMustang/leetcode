package main

// 输入一棵二叉树的根节点，求该树的深度。
// 从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func maxDepth(root *TreeNode) int {
// 	if root == nil {
// 		return 0
// 	}
// 	res, size := 0, 1
// 	queue := []*TreeNode{root}
// 	for size != 0 {
// 		res++
// 		for i := 0; i < size; i++ {
// 			node := queue[0]
// 			if node.Left != nil {
// 				queue = append(queue, node.Left)
// 			}
// 			if node.Right != nil {
// 				queue = append(queue, node.Right)
// 			}
// 			queue = queue[1:]
// 		}
// 		size = len(queue)
// 	}
// 	return res
// }
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
