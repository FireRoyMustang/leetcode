package main

// 给定一个二叉搜索树, 找到该树中两个指定节点的最近公共祖先。

// 百度百科中最近公共祖先的定义为：
// “对于有根树 T 的两个结点 p、q，最近公共祖先表示为一个结点 x，满足 x 是 p、q 的祖先且 x 的深度
// 尽可能大（一个节点也可以是它自己的祖先）。”

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

// 迭代
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if root.Val > p.Val && root.Val > q.Val {
			root = root.Right
		} else {
			break
		}
	}

	return root
}

// 递归
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	if root == nil {
// 		return nil
// 	}

// 	if p.Val < root.Val && q.Val < root.Val {
// 		return lowestCommonAncestor(root.Left, p, q)
// 	}
// 	if root.Val > p.Val && root.Val > q.Val {
// 		return lowestCommonAncestor(root.Right, p, q)
// 	}
// 	return root
// }
