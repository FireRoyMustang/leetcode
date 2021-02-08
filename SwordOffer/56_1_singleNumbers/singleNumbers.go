package main

import "fmt"

// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
// 请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

func main() {
	// nums := []int{4, 1, 4, 6}
	nums := []int{4, 1, 4, 6}
	fmt.Println(singleNumbers(nums))
}

func singleNumbers(nums []int) []int {
	tmp := 0
	for i := 0; i < len(nums); i++ {
		tmp ^= nums[i]
	}
	div := 1
	for tmp&div == 0 {
		div <<= 1
	}
	a, b := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i]&div == 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}
