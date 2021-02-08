package main

import "fmt"

// 输入一个整型数组，数组中的一个或连续多个整数组成一个子数组。求所有子数组的和的最大值。
// 要求时间复杂度为O(n)。
func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		res = max(nums[i], res)
	}
	return res
}

// func maxSubArray(nums []int) int {
// 	length := len(nums)
// 	if length == 1 {
// 		return nums[0]
// 	}
// 	res, tmp := -1<<31, 0

// 	for _, num := range nums {
// 		if num < 0 {
// 			if res < 0 {
// 				res = max(res, num)
// 			}
// 			if tmp+num < 0 {
// 				tmp = 0
// 			} else {
// 				tmp += num
// 			}
// 		} else {
// 			tmp += num
// 			res = max(res, tmp)
// 		}
// 	}
// 	return res
// }
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
