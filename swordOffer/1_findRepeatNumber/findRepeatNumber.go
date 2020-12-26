package main

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
// 数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。
// 请找出数组中任意一个重复的数字。

func main() {

}
func findRepeatNumber(nums []int) int {
	length := len(nums)
	arr := make([]int, length)
	for i := 0; i < length; i++ {
		arr[nums[i]]++
		if arr[nums[i]] >= 2 {
			return nums[i]
		}
	}
	return -1

}
