package main

import "fmt"

// 给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，
// 使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums) // 0,0,1,1,2,2
	fmt.Println(nums)
}

func sortColors(nums []int) {
	left, right := 0, len(nums)-1 // left表示0左边界，right表示1左边界
	for i := 0; i <= right; i++ {
		for i <= right && nums[i] == 2 {
			nums[i], nums[right] = nums[right], nums[i]
			right--
		}
		if nums[i] == 0 {
			nums[i], nums[left] = nums[left], nums[i]
			left++
		}
	}
}
