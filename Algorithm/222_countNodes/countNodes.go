package main

import "sort"

// 给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。

// 完全二叉树 的定义如下：在完全二叉树中，除了最底层节点可能没填满外，
// 其余每层节点数都达到最大值，并且最下面一层的节点都集中在该层最左边的若干位置。
// 若最底层为第 h 层，则该层包含 1~ 2^h 个节点。
// 树中节点的数目范围是[0, 5 * 104]
// 0 <= Node.val <= 5 * 104
// 题目数据保证输入的树是 完全二叉树
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
	for node := root; node.Left != nil; node = node.Left {
		level++
	}
	return sort.Search(1<<(level+1), func(k int) bool {
		if k <= 1<<level {
			return false
		}
		bits := 1 << (level - 1)
		node := root
		for node != nil && bits > 0 {
			if bits&k == 0 {
				node = node.Left
			} else {
				node = node.Right
			}
			bits >>= 1
		}
		return node == nil
	}) - 1
}

// func countNodes(root *TreeNode) int {
// 	res := 0
// 	if root == nil {
// 		return res
// 	}
// 	queue := []*TreeNode{root}
// 	size := len(queue)
// 	for size != 0 {
// 		for i := 0; i < size; i++ {
// 			node := queue[i]
// 			if node.Left == nil {
// 				return res
// 			}
// 			queue = append(queue, node.Left)
// 			res++
// 			if node.Right == nil {
// 				return res
// 			}
// 			queue = append(queue, node.Right)
// 			res++
// 		}
// 		queue = queue[size:]
// 		size = len(queue)
// 	}
// 	return res
// }
