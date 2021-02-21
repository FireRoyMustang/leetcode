package main

import "fmt"

// 给定一个数组 nums 和滑动窗口的大小 k，请找出所有滑动窗口里的最大值。

func main() {
	// fmt.Println(maxSlidingWindow([]int{5, 3, 4}, 1))
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length == 0 {
		return []int{}
	}
	queue, res := make([]int, 0), make([]int, 0)
	for i := 0; i < k; i++ {
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	for i := k; i < length; i++ {
		// pop
		res = append(res, queue[0])
		if nums[i-k] == queue[0] {
			queue = queue[1:]
		}
		// push
		for len(queue) > 0 && queue[len(queue)-1] < nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, nums[i])
	}
	res = append(res, queue[0])
	return res
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	length := len(nums)
// 	if length == 0 {
// 		return []int{}
// 	}
// 	if length <= k {
// 		return []int{max(nums, length)}
// 	}
// 	res := make([]int, length-k+1)
// 	for i := 0; i < length-k+1; i++ {
// 		res[i] = max(nums[i:i+k], k)
// 	}

// 	return res
// }
// func max(nums []int, k int) int {
// 	res := nums[0]
// 	for i := 1; i < k; i++ {
// 		if res < nums[i] {
// 			res = nums[i]
// 		}
// 	}
// 	return res
// }
