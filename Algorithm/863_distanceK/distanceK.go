package main

// 给定的树是非空的。
// 树上的每个结点都具有唯一的值 0 <= node.val <= 500 。
// 目标结点 target 是树上的结点。
// 0 <= K <= 1000.

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	mp  map[*TreeNode]*TreeNode
	res []int
)

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	mp = make(map[*TreeNode]*TreeNode)
	res = make([]int, 0)
	makeMap(root, nil)
	// queue := make([]*TreeNode, 0)
	// queue = append(queue, target)
	queue := []*TreeNode{target}
	size := len(queue)
	seen := make(map[*TreeNode]bool)
	for K > 0 && size != 0 {
		for size != 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if seen[node] {
				continue
			}
			seen[node] = true
			if node.Left != nil && !seen[node.Left] {
				queue = append(queue, node.Left)
			}
			if node.Right != nil && !seen[node.Right] {
				queue = append(queue, node.Right)
			}
			if mp[node] != nil && !seen[mp[node]] {
				queue = append(queue, mp[node])
			}
		}
		size = len(queue)
		K--
	}
	for _, node := range queue {
		res = append(res, node.Val)
	}
	return res
}
func makeMap(root, parent *TreeNode) {
	if root == nil {
		return
	}
	mp[root] = parent
	makeMap(root.Left, root)
	makeMap(root.Right, root)
}
