package main

// 输入一棵二叉搜索树，将该二叉搜索树转换成一个排序的循环双向链表。
// 要求不能创建任何新的节点，只能调整树中节点指针的指向。
// 为了让您更好地理解问题，以下面的二叉搜索树为例：
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	head *Node
	pre  *Node
)

func treeToDoublyList(root *TreeNode) *Node {
	if root == nil {
		return nil
	}
	dfs(root)
	pre.Right = head
	head.Left = pre
	return head
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	node := &Node{root.Val, nil, nil}
	dfs(root.Left)
	if pre == nil {
		head = node
	} else {
		node.Left = pre
		pre.Right = node
	}
	pre = node
	dfs(root.Right)
}
