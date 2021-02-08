package main

// 给定一棵二叉搜索树，请找出其中第k大的节点。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthLargest(root *TreeNode, k int) int {
	return dfs(root, &k)
}
func dfs(root *TreeNode, k *int) int {
	if root == nil {
		return -1
	}
	res := dfs(root.Right, k)
	if res != -1 {
		return res
	}
	if *k == 1 {
		return root.Val
	}
	*k--
	res = dfs(root.Left, k)
	if res != -1 {
		return res
	}
	return -1
}
