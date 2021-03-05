package main

import "fmt"

// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，
// 使得所有奇数位于数组的前半部分，所有偶数位于数组的后半部分。

func main() {
	testSlice := []int{1, 2, 3, 4}
	fmt.Println(exchange(testSlice))
}
func exchange(nums []int) []int {
	length := len(nums)
	ptL, ptR := 0, length-1
	for ptL < ptR {
		for ptL < ptR && nums[ptL]%2 == 1 {
			ptL++
		}
		for ptL < ptR && nums[ptR]%2 == 0 {
			ptR--
		}
		nums[ptL], nums[ptR] = nums[ptR], nums[ptL]
	}
	return nums
}
