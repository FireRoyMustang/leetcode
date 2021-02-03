package main

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
	node.Left = pre
	if pre == nil {
		head = node
	} else {
		node.Left = pre
		pre.Right = node
	}
	pre = node
	dfs(root.Right)
}
