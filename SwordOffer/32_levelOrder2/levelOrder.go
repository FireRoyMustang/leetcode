package main

// 从上到下按层打印二叉树，同一层的节点按从左到右的顺序打印，每一层打印到一行。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		size := len(queue)
		floor := make([]int, size)
		index := 0
		res = append(res, floor)
		for size != 0 {
			node := queue[0]
			queue = queue[1:]
			floor[index] = node.Val
			index++
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			size--
		}
	}
	return res
}
