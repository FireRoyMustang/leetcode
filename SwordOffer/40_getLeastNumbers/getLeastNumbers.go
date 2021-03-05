package main

import (
	"container/heap"
	"fmt"
)

// 输入整数数组 arr ，找出其中最小的 k 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，
// 则最小的4个数字是1、2、3、4。

func main() {
	// test := []int{4, 5, 1, 6, 2, 7, 3, 8}
	test := []int{0, 0, 0, 2, 0, 5}
	fmt.Println(getLeastNumbers(test, 0))
}

func getLeastNumbers(arr []int, k int) []int {
	length := len(arr)
	if k == 0 {
		return []int{}
	}
	if length == k {
		return arr
	}
	res := make([]int, k)
	maxHeap := &IntHeap{}
	for i := 0; i < k; i++ {
		heap.Push(maxHeap, arr[i])

	}
	for i := k; i < length; i++ {
		if (*maxHeap)[0] > arr[i] {
			heap.Pop(maxHeap)
			heap.Push(maxHeap, arr[i])
		}

	}
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(maxHeap).(int)
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(i interface{}) {
	*h = append(*h, i.(int))
}
func (h *IntHeap) Pop() interface{} {
	length := h.Len()
	popped := (*h)[length-1]
	*h = (*h)[:length-1]
	return popped
}
