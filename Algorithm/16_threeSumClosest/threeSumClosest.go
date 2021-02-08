package main

import (
	"fmt"
	"sort"
)

// 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。
// 找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

func main() {
	// fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}

func threeSumClosest(nums []int, target int) int {
	length := len(nums)
	if length <= 2 {
		return -1
	}
	if length == 3 {
		return nums[0] + nums[1] + nums[2]
	}
	sort.Ints(nums)
	diff := nums[0] + nums[1] + nums[2] - target
	for k := 0; k < length-2; k++ {
		lptr := k + 1
		rptr := length - 1
		for lptr < rptr {
			sum := nums[k] + nums[lptr] + nums[rptr]
			tmp := sum - target
			if tmp == 0 {
				return target
			} else if tmp < 0 {
				lptr++
			} else {
				rptr--
			}
			if abs(tmp) < abs(diff) {
				diff = tmp
			}
		}
	}
	return target + diff
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
