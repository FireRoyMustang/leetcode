package main

import "fmt"

// 给定一个非空且只包含非负数的整数数组 nums，数组的度的定义是指数组里任一元素出现频数的最大值。
// 你的任务是在 nums 中找到与 nums 拥有相同大小的度的最短连续子数组，返回其长度。

func main() {
	// fmt.Println(findShortestSubArray(([]int{1, 2, 2, 3, 1})))
	fmt.Println(findShortestSubArray(([]int{1, 2, 2, 3, 1, 4, 2})))
}

type Helper struct {
	Start int
	Freq  int
}

func findShortestSubArray(nums []int) int {
	mp := make(map[int]*Helper)
	res, maxFreq := 0, 0
	for index, num := range nums {
		helper, has := mp[num]
		if !has {
			mp[num] = &Helper{index, 0}
			helper = mp[num]
		}
		helper.Freq++
		if helper.Freq > maxFreq {
			maxFreq = helper.Freq
			res = index - helper.Start + 1
		} else if helper.Freq == maxFreq {
			res = min(index-helper.Start+1, res)
		}
	}
	return res
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
