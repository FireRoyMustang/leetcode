package tools

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func GenerateTreeNode(nodeVal int) *TreeNode {
	if nodeVal == -1 {
		return nil
	}
	return &TreeNode{nodeVal, nil, nil}
}

// 层序遍历二叉树
func LevelOrder(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 1)
	queue[0] = root
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			node := queue[0]
			res = append(res, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
			size--
		}
	}
	return res
}

// 根据数组，生成树
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
