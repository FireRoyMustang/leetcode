package main

// 输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
// 如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。

import "fmt"

func main() {
	// test := []int{1, 6, 3, 2, 5}
	// test := []int{4, 8, 6, 12, 16, 14, 10}
	// test := []int{1, 2, 5, 10, 6, 9, 4, 3}
	test := []int{1, 2, 3, 4}
	fmt.Println(verifyPostorder(test))
}

func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length <= 1 {
		return true
	}
	i := 0
	for i = 0; i < length; i++ {
		if postorder[i] > postorder[length-1] {
			break
		}
	}
	mid := i
	for ; i < length; i++ {
		if postorder[i] < postorder[length-1] {
			return false
		}
	}
	if mid == length {
		return verifyPostorder(postorder[:mid-1])
	}
	return verifyPostorder(postorder[:mid]) && verifyPostorder(postorder[mid:length-1])
}
