package main

import "fmt"

// 给定一个只包含正整数的非空数组。是否可以将这个数组分割成两个子集，使得两个子集的元素和相等。

// 注意:
// 每个数组中的元素不会超过 100
// 数组的大小不会超过 200

func main() {
	fmt.Println(canPartition([]int{3, 3, 3, 4, 5}))
}

func canPartition(nums []int) bool {
	length := len(nums)
	if length <= 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	// dp[i][j]
	dp[0] = true
	for i := 0; i < length; i++ {
		if nums[i] > sum {
			continue
		}
		if dp[sum-nums[i]] {
			return true
		}
		for v := sum - 1; v >= nums[i]; v-- {
			dp[v] = dp[v-nums[i]] || dp[v]
		}
	}
	return dp[sum]
}
