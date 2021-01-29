package main

// 请实现一个函数按照之字形顺序打印二叉树，即第一行按照从左到右的顺序打印，
// 第二层按照从右到左的顺序打印，第三行再按照从左到右的顺序打印，其他行以此类推。
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
	flag := true
	index := 0
	for len(queue) != 0 {
		size := len(queue)
		floor := make([]int, size)
		if flag {
			index = 0
		} else {
			index = size - 1
		}
		res = append(res, floor)
		for size != 0 {
			node := queue[0]
			floor[index] = node.Val
			if flag {
				index++
			} else {
				index--
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			size--
		}
		flag = !flag
	}
	return res
}
