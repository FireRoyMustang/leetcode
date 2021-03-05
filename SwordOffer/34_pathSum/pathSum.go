package main

import (
	"fmt"
)

// 输入一棵二叉树和一个整数，打印出二叉树中节点值的和为输入整数的所有路径。
// 从树的根节点开始往下一直到叶节点所经过的节点形成一条路径。
// https://blog.csdn.net/sgsgy5/article/details/81590184
// append
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	test := []int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}
	tree := GenerateTree(test)
	fmt.Println(pathSum(tree, 22))

}

func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int = make([][]int, 0)
	if root == nil {
		return res
	}
	backTrace(root, []int{}, sum, &res)
	return res
}

func backTrace(root *TreeNode, path []int, sum int, res *[][]int) {
	if root == nil {
		return
	}
	sum -= root.Val
	path = append(path, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		copyPath := make([]int, len(path))
		copy(copyPath, path)
		*res = append(*res, copyPath)
	}
	if root.Left != nil {
		backTrace(root.Left, path, sum, res)
	}
	if root.Right != nil {
		backTrace(root.Right, path, sum, res)
	}
}

func GenerateTreeNode(nodeVal int) *TreeNode {
	if nodeVal == -1 {
		return nil
	}
	return &TreeNode{nodeVal, nil, nil}
}

func GenerateTree(treeNums []int) *TreeNode {
	length := len(treeNums)
	if length == 0 {
		return nil
	}
	treeNumsIndex := 1
	root := GenerateTreeNode(treeNums[0])
	queue := []*TreeNode{root}
	for treeNumsIndex < length {
		node := queue[0]
		queue = queue[1:]
		node.Left = GenerateTreeNode(treeNums[treeNumsIndex])
		treeNumsIndex++
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if treeNumsIndex >= length {
			break
		}

		node.Right = GenerateTreeNode(treeNums[treeNumsIndex])
		treeNumsIndex++
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return root
}
