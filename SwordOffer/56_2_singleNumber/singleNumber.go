package main

import "fmt"

// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

func main() {
	test := []int{3, 4, 3, 3}
	fmt.Println(singleNumber(test))
}

func singleNumber(nums []int) int {
	ones, twos := 0, 0
	for _, num := range nums {
		ones = ones ^ num&^twos
		twos = twos ^ num&^ones
	}
	return ones
}
