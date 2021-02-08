package main

import "fmt"

// 当 A 的子数组 A[i], A[i+1], ..., A[j] 满足下列条件时，我们称其为湍流子数组：

// 若 i <= k < j，当 k 为奇数时， A[k] > A[k+1]，且当 k 为偶数时，A[k] < A[k+1]；
// 或 若 i <= k < j，当 k 为偶数时，A[k] > A[k+1] ，且当 k 为奇数时， A[k] < A[k+1]。
// 也就是说，如果比较符号在子数组中的每个相邻元素对之间翻转，则该子数组是湍流子数组。

// 返回 A 的最大湍流子数组的长度。

func main() {
	testSlices := [][]int{{9, 4, 2, 10, 7, 8, 8, 1, 9},
		{4, 8, 12, 16},
		{100}}
	for _, testSlice := range testSlices {
		fmt.Println(maxTurbulenceSize(testSlice))
	}
}

// func maxTurbulenceSize(arr []int) int {
// 	n := len(arr)
// 	ans := 1
// 	left, right := 0, 0
// 	for right < n-1 {
// 		if left == right {
// 			if arr[left] == arr[left+1] {
// 				left++
// 			}
// 			right++
// 		} else {
// 			if arr[right-1] < arr[right] && arr[right] > arr[right+1] {
// 				right++
// 			} else if arr[right-1] > arr[right] && arr[right] < arr[right+1] {
// 				right++
// 			} else {
// 				left = right
// 			}
// 		}
// 		ans = max(ans, right-left+1)
// 	}
// 	return ans
// }

func maxTurbulenceSize(arr []int) int {
	length := len(arr)
	if length <= 1 {
		return length
	}
	res := 1
	flag, pre, start := 0, arr[0], 0 //0代表无状态，1代表小于，2代表大于
	for i := 1; i < length; i++ {
		if arr[i] != pre {
			if flag == 0 {
				res = max(i-start+1, res)
			} else if flag == 1 {
				if arr[i] < pre {
					start = i - 1
				} else {
					res = max(i-start+1, res)
				}
			} else {
				if arr[i] > pre {
					start = i - 1
				} else {
					res = max(i-start+1, res)
				}
			}
			if arr[i] < pre {
				flag = 1
			} else {
				flag = 2
			}
		} else {
			flag = 0
			start = i
		}
		pre = arr[i]
	}
	return res
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
