package main

// 输入一棵二叉树的根节点，判断该树是不是平衡二叉树。
// 如果某二叉树中任意节点的左右子树的深度相差不超过1，那么它就是一棵平衡二叉树。
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return helper(root) != -1
}

func helper(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := helper(root.Left)
	if left == -1 {
		return -1
	}
	right := helper(root.Right)
	if right == -1 {
		return -1
	}
	if abs(left-right) <= 1 {
		return max(left, right) + 1
	}
	return -1
}
func abs(i int) int {
	if i > 0 {
		return i
	}
	return -i
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
