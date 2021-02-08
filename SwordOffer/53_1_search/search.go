package main

import "fmt"

// 统计一个数字在排序数组中出现的次数。

func main() {
	test := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(search(test, 8))
	// fmt.Println(search(test, 6))
}

// func search(nums []int, target int) int {
// 	res := 0
// 	for _, num := range nums {
// 		if num == target {
// 			res++
// 		}
// 	}
// 	return res
// }

// 二分法
func search(nums []int, target int) int {
	return helper(nums, target) - helper(nums, target-1)
}
func helper(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
