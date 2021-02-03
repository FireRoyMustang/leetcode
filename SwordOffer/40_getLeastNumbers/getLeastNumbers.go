package main

import "container/list"

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。

func main() {

}

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)
	if length < k {
		return []int{}
	}
	if length == k {
		return arr
	}
	res := make([]int, k)
	var dequeue list.List
	for _, num := range arr {
		if dequeue.Len() < k {
			val, _ := dequeue.Back().Value.(int)
			if num < val {

			}
		}
	}

	return res
}
