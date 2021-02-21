package main

import "fmt"

// 输入一个递增排序的数组和一个数字s，在数组中查找两个数，
// 使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

func main() {
	test := []int{2, 7, 11, 15}
	fmt.Println(twoSum(test, 9))
}

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{nums[left], nums[right]}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return []int{nums[left], nums[right]}
}

// func twoSum(nums []int, target int) []int {
// 	left, right := 0, len(nums)-1
// 	for left < right {
// 		if nums[left]+nums[right] == target {
// 			return []int{nums[left], nums[right]}
// 		} else if nums[left]+nums[right] > target {
// 			right--
// 		} else {
// 			left++
// 		}
// 	}
// 	return []int{nums[left], nums[right]}
// }
