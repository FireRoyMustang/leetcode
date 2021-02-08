package main

import (
	"fmt"
	"sort"
)

// 给定长度为 2n 的整数数组 nums ，你的任务是将这些数分成 n 对,
// 例如 (a1, b1), (a2, b2), ..., (an, bn) ，使得从 1 到 n 的 min(ai, bi) 总和最大

func main() {
	testSlice := []int{1, 4, 3, 2}
	// testSlice := []int{6, 2, 6, 5, 1, 2}
	fmt.Println(arrayPairSum(testSlice))
}

func arrayPairSum(nums []int) int {
	sort.Ints(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
