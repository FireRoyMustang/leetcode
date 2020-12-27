package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	length := len(preorder)
	if length == 0 {
		return nil
	}

	m := make(map[int]int, length)
	for index, num := range inorder {
		m[num] = index
	}
	root := helper(preorder, 0, length-1, inorder, 0, length-1, m)
	return root

}
func helper(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int,
	indexMap map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}
	rootVal := preorder[preStart]
	rootIndex := indexMap[rootVal]
	root := &TreeNode{
		Val:   rootVal,
		Left:  nil,
		Right: nil,
	}
	length := rootIndex - inStart
	root.Left = helper(preorder, preStart+1, preStart+length,
		inorder, inStart, rootIndex-1, indexMap)
	root.Right = helper(preorder, preStart+length+1, preEnd,
		inorder, rootIndex+1, inEnd, indexMap)
	return root
}
