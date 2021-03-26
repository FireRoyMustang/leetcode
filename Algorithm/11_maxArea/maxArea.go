package main
// 给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
// 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。
// 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

// 说明：你不能倾斜容器。

import "fmt"

func main() {
	// testSlice := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// testSlice := []int{1, 1}
	testSlice := []int{2, 3, 10, 5, 7, 8, 9}
	fmt.Println(maxArea(testSlice))
}
func maxArea(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}
	ans, left, right, area := 0, 0, length-1, 0
	for left < right {
		area = min(height[left], height[right]) * (right - left)
		ans = max(ans, area)
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
