package main

import "fmt"

// 我们把只包含质因子 2、3 和 5 的数称作丑数（Ugly Number）。求按从小到大的顺序的第 n 个丑数。

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(nthUglyNumber(i))
	}
}

func nthUglyNumber(n int) int {
	ugArr := make([]int, n)
	ugArr[0] = 1
	ptr2, ptr3, ptr5 := 0, 0, 0
	for i := 1; i < n; i++ {
		ug2 := ugArr[ptr2] * 2
		ug3 := ugArr[ptr3] * 3
		ug5 := ugArr[ptr5] * 5
		ugArr[i] = min(ug2, ug3, ug5)
		if ugArr[i] == ug2 {
			ptr2++
		}
		if ugArr[i] == ug3 {
			ptr3++
		}
		if ugArr[i] == ug5 {
			ptr5++
		}
	}
	return ugArr[n-1]
}
func min(i, j, k int) int {
	return min2(min2(i, j), k)
}
func min2(i, j int) int {
	if i < j {
		return i
	}
	return j
}
