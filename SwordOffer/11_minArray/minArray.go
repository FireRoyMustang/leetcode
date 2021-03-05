package main

import "fmt"

// 把一个数组最开始的若干个元素搬到数组的末尾，我们称之为数组的旋转。
// 输入一个递增排序的数组的一个旋转，输出旋转数组的最小元素。
// 例如，数组 [3,4,5,1,2] 为 [1,2,3,4,5] 的一个旋转，该数组的最小值为1。

func main() {
	// fmt.Println(minArray([]int{3, 4, 5, 1, 2})) // 1
	// fmt.Println(minArray([]int{2, 2, 2, 0, 1})) // 0
	// fmt.Println(minArray([]int{1, 2, 3, 4, 5})) // 0
	fmt.Println(minArray([]int{3, 3, 1, 3})) // 1
	// fmt.Println(minArray([]int{3, 4, 5, 1}))    // 1
}
func minArray(numbers []int) int {
	length := len(numbers)
	left, right := 0, length-1
	mid := 0
	for left < right {
		mid = (left + right) >> 1
		if numbers[mid] < numbers[right] {
			right = mid
		} else if numbers[mid] > numbers[right] {
			left = mid + 1
		} else {
			right--
			// right = mid //wrong
		}
	}
	return numbers[left]
}
