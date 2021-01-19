package main

import "fmt"

// 给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

func main() {
	// test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// test := []int{}
	test := []int{1, 2}
	rotate(test, 11)
	fmt.Println(test)
}
func rotate(nums []int, k int) {
	length := len(nums)
	if length <= 1 || k%length == 0 {
		return
	}
	if k > length {
		k %= length
	}
	tmp := make([]int, length-k)
	copy(tmp, nums[:length-k])
	copy(nums, nums[length-k:])
	copy(nums[k:], tmp)
}
