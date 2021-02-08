package main

import "fmt"

func main() {
	test := []int{2, 8, 7, 4, 6, 3}
	fmt.Println(mergeSort(test))

}

func mergeSort(arr []int) []int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := len(arr) >> 1
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(arr1 []int, arr2 []int) []int {
	len1 := len(arr1)
	len2 := len(arr2)
	res := make([]int, len1+len2)
	ptr, ptr1, ptr2 := 0, 0, 0
	for ptr1 < len1 && ptr2 < len2 {
		if arr1[ptr1] < arr2[ptr2] {
			res[ptr] = arr1[ptr1]
			ptr1++
		} else {
			res[ptr] = arr2[ptr2]
			ptr2++
		}
		ptr++
	}
	for ; ptr2 < len2; ptr2++ {
		res[ptr] = arr2[ptr2]
		ptr++
	}
	for ; ptr1 < len1; ptr1++ {
		res[ptr] = arr1[ptr1]
		ptr++

	}
	return res
}
