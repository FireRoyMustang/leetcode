package main

import "fmt"

// 一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
// 在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

func main() {
	// testSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 9}
	testSlice := []int{0}
	fmt.Println(missingNumber(testSlice))
}

func missingNumber(nums []int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] <= mid {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}

// func missingNumber(nums []int) int {
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] != i {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }
