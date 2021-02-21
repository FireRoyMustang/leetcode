package main

import "fmt"

// 给你一个整数数组 nums ，和一个表示限制的整数 limit，请你返回最长连续子数组的长度，
// 该子数组中的任意两个元素之间的绝对差必须小于或者等于 limit 。
// 如果不存在满足条件的子数组，则返回 0 。

func main() {
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4))
}

func longestSubarray(nums []int, limit int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}
	res, start := 0, 0
	minQueue := make([]int, 0)
	maxQueue := make([]int, 0)
	for index, num := range nums {
		for len(minQueue) > 0 && minQueue[len(minQueue)-1] > num {
			minQueue = minQueue[:len(minQueue)-1]
		}
		minQueue = append(minQueue, num)
		for len(maxQueue) > 0 && maxQueue[len(maxQueue)-1] < num {
			maxQueue = maxQueue[:len(maxQueue)-1]
		}
		maxQueue = append(maxQueue, num)
		diff := maxQueue[0] - minQueue[0]
		if diff <= limit {
			res = max(res, index-start+1)
		} else {
			for diff > limit { // pop
				if nums[start] == maxQueue[0] {
					maxQueue = maxQueue[1:]
				}
				if nums[start] == minQueue[0] {
					minQueue = minQueue[1:]
				}
				diff = maxQueue[0] - minQueue[0]
				start++
			}
		}

	}
	return res
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
