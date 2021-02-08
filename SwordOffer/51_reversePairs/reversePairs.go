package main

import "fmt"

// 在数组中的两个数字，如果前面一个数字大于后面的数字，
// 则这两个数字组成一个逆序对。输入一个数组，求出这个数组中的逆序对的总数。
func main() {
	testSlice := []int{7, 5, 6, 4}
	fmt.Println(reversePairs(testSlice))
}

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}
func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := left + (right-left)>>1
	cnt := mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	ptr, i, j, tmp := 0, left, mid+1, make([]int, right-left+1)

	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp[ptr] = nums[i]
			i++
			cnt += j - (mid + 1)
		} else {
			tmp[ptr] = nums[j]
			j++
		}
		ptr++
	}
	for i <= mid {
		tmp[ptr] = nums[i]
		// cnt += j - (mid + 1)
		cnt += right - mid
		i++
		ptr++
	}
	for j <= right {
		tmp[ptr] = nums[j]
		j++
		ptr++
	}
	copy(nums[left:right+1], tmp)
	// for i := 0; i < len(tmp); i++ {
	// 	nums[left+i] = tmp[i]
	// }
	return cnt
}
