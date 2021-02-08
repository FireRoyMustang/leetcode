package main

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
