package main

import (
	"fmt"
	"sort"
)

// 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，
// 使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

func main() {
	// test := []int{-1, 0, 1, 2, -1, -4}
	test := []int{0, 0, 0}
	fmt.Println(threeSum(test))
}
func threeSum(nums []int) [][]int {
	length := len(nums)
	if length <= 2 {
		return nil
	}
	ans := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < length-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lptr := i + 1
		rptr := length - 1
		for lptr < rptr {
			sum := nums[i] + nums[lptr] + nums[rptr]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[lptr], nums[rptr]})
				lptr++
				for lptr < rptr && nums[lptr] == nums[lptr-1] {
					lptr++
				}
				rptr--
				for lptr < rptr && nums[rptr] == nums[rptr+1] {
					rptr--
				}
			} else if sum < 0 {
				lptr++
				for lptr < rptr && nums[lptr] == nums[lptr-1] {
					lptr++
				}
			} else {
				rptr--
				for lptr < rptr && nums[rptr] == nums[rptr+1] {
					rptr--
				}
			}
		}
	}
	return ans
}
