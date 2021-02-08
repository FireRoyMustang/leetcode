package main

import "sort"

func nextPermutation(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	a := length - 2
	for ; a >= 0 && nums[a] >= nums[a+1]; a-- {
	}
	if a == -1 {
		sort.Ints(nums)
		return
	}
	b := length - 1
	for ; b >= 0 && nums[b] <= nums[a]; b-- {
	}
	swap(a, b, nums)
	for i, j := a+1, length-1; i < j; i, j = i+1, j-1 {
		swap(i, j, nums)
	}
}
func swap(i, j int, nums []int) {
	nums[i], nums[j] = nums[j], nums[i]
}
