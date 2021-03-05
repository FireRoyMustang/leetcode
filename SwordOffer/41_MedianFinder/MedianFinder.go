package main

import (
	"container/heap"
	"fmt"
)

// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是
// 所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是
// 所有数值排序之后中间两个数的平均值。

// 例如，

// [2,3,4] 的中位数是 3

// [2,3] 的中位数是 (2 + 3) / 2 = 2.5

// 设计一个支持以下两种操作的数据结构：

func main() {
	// test := []int{2, 8, 6}
	test := []int{1, 2, 3, 6}
	mid := Constructor()
	for i := 0; i < len(test); i++ {
		mid.AddNum(test[i])
		fmt.Println(mid.FindMedian())
	}
}

type MedianFinder struct {
	maxHeap *IntHeap
	minHeap *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{&IntHeap{true, make([]int, 0)}, &IntHeap{false, make([]int, 0)}}
}

func (this *MedianFinder) AddNum(num int) {
	minHeapLen := len(this.minHeap.elements)
	if minHeapLen == 0 {
		heap.Push(this.minHeap, num)
		return
	}
	maxHeapLen := len(this.maxHeap.elements)
	if minHeapLen == maxHeapLen {
		if num >= this.maxHeap.elements[0] {
			heap.Push(this.minHeap, num)
		} else {
			heap.Push(this.maxHeap, num)
			heap.Push(this.minHeap, heap.Pop(this.maxHeap))
		}
	} else {
		if num > this.minHeap.elements[0] {
			heap.Push(this.minHeap, num)
			heap.Push(this.maxHeap, heap.Pop(this.minHeap))
		} else {
			heap.Push(this.maxHeap, num)
		}
	}

}

func (this *MedianFinder) FindMedian() float64 {
	minHeapLen := this.minHeap.Len()
	maxHeapLen := this.maxHeap.Len()
	if minHeapLen == maxHeapLen {
		res := 0
		res += this.minHeap.elements[0]
		res += this.maxHeap.elements[0]
		return float64(res) / 2
	}
	return float64(this.minHeap.elements[0])
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

type IntHeap struct {
	flag     bool // false为小根堆，true为大根堆
	elements []int
}

func (h IntHeap) Len() int { return len(h.elements) }
func (h IntHeap) Less(i, j int) bool {
	if h.flag {
		return h.elements[i] > h.elements[j]
	}
	return h.elements[i] < h.elements[j]
}
func (h IntHeap) Swap(i, j int) { h.elements[i], h.elements[j] = h.elements[j], h.elements[i] }

func (h *IntHeap) Push(i interface{}) {
	h.elements = append(h.elements, i.(int))
}
func (h *IntHeap) Pop() interface{} {
	length := h.Len()
	popped := h.elements[length-1]
	h.elements = h.elements[:length-1]
	return popped
}
