package main

import (
	"fmt"
)

func main() {
	test := []int{2, 8, 7, 4, 6, 3}
	QuickSort(test)
	fmt.Println(test)
}

func QuickSort(arr []int) {
	quick(arr, 0, len(arr)-1)
}

func quick(arr []int, left, right int) {
	if left >= right {
		return
	}
	pivot, ptr, i, j := arr[left], left, left, right
	for i < j {
		for j > ptr && arr[j] >= pivot {
			j--
		}
		if j > ptr {
			arr[ptr] = arr[j]
			ptr = j
		}
		for i < ptr && arr[i] <= pivot {
			i++
		}
		if i < ptr {
			arr[ptr] = arr[i]
			ptr = i
		}
	}
	arr[ptr] = pivot

	quick(arr, left, ptr-1)
	quick(arr, ptr+1, right)
}

// 第一种写法
func quickSort(values []int, left, right int) {
	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}
	}
	values[p] = temp
	if p-left > 1 {
		quickSort(values, left, p-1)
	}
	if right-p > 1 {
		quickSort(values, p+1, right)
	}
}

func QuickSort1(values []int) {
	if len(values) <= 1 {
		return
	}
	quickSort(values, 0, len(values)-1)
}
