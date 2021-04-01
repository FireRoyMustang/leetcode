package main

import (
	"fmt"
	"sort"
)

// 给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。

// 解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{5, 5, 5}))
	fmt.Println(subsetsWithDup([]int{5, 5, 5, 5}))
}

func subsetsWithDup(nums []int) [][]int {
	size := len(nums)
	if size == 0 {
		return [][]int{[]int{}}
	}
	res := [][]int{}
	sort.Ints(nums)
	dfs(nums, 0, size, &res, []int{}, false)
	return res
}

func dfs(nums []int, index, size int, res *[][]int, ele []int, choose bool) {
	if index == size {
		*res = append(*res, ele)
		return
	}
	dfs(nums, index+1, size, res, ele, false)
	// 如果和上一个元素相同并且没有选择上一个元素则不需要加入该元素
	if index > 0 && nums[index] == nums[index-1] && !choose {
		return
	}
	eleCopy := make([]int, len(ele))
	copy(eleCopy, ele)
	eleCopy = append(eleCopy, nums[index])
	dfs(nums, index+1, size, res, eleCopy, true)
}
