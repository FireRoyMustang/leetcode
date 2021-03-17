package main

import "fmt"

// 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
// n == height.length
// 0 <= n <= 3 * 104
// 0 <= height[i] <= 105
func main() {
	fmt.Println(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})) // 6
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))                   // 9
}

func trap(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	maxLeft, maxRight := 0, 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				res += (maxLeft - height[left])
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				res += (maxRight - height[right])
			}
			right--
		}
	}
	return res
}

// func trap(height []int) int {
// 	n := len(height)
// 	if n <= 1 {
// 		return 0
// 	}
// 	res := 0
// 	maxLeft, maxRight := 0, make([]int, n)
// 	maxRight[n-1] = height[n-1]
// 	for i := n - 2; i >= 0; i-- {
// 		maxRight[i] = max(maxRight[i+1], height[i])
// 	}
// 	for i := 0; i < n; i++ {
// 		maxLeft = max(maxLeft, height[i])
// 		res += min(maxLeft, maxRight[i]) - height[i]
// 	}
// 	return res
// }

// func trap(height []int) int {
// 	n := len(height)
// 	if n <= 1 {
// 		return 0
// 	}
// 	res := 0
// 	for i := 0; i < n; i++ {
// 		maxLeft, maxRight := 0, 0
// 		for j := i; j < n; j++ {
// 			maxRight = max(maxRight, height[j])
// 		}
// 		for j := i; j >= 0; j-- {
// 			maxLeft = max(maxLeft, height[j])
// 		}
// 		res += min(maxLeft, maxRight) - height[i]

// 	}
// 	return res
// }
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
