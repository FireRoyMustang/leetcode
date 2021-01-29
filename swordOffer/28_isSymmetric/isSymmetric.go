package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{1, nil, nil}
	root.Left = &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}
	root.Right = &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}
	isSymmetric(root)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return recur(root.Left, root.Right)
}
func recur(tree1, tree2 *TreeNode) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil {
		return false
	}
	return tree1.Val == tree2.Val && recur(tree1.Left, tree2.Right) && recur(tree1.Right, tree2.Left)
}
