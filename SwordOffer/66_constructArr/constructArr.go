package main

import "fmt"

// 给定一个数组 A[0,1,…,n-1]，请构建一个数组 B[0,1,…,n-1]，
// 其中 B[i] 的值是数组 A 中除了下标 i 以外的元素的积,
// 即 B[i]=A[0]×A[1]×…×A[i-1]×A[i+1]×…×A[n-1]。
// 不能使用除法。

// 所有元素乘积之和不会溢出 32 位整数
// a.length <= 100000

func main() {
	fmt.Println(constructArr([]int{1, 2, 3, 4, 5}))
	fmt.Println(constructArr([]int{}))
}
func constructArr(a []int) []int {
	res := make([]int, len(a))
	if len(a) == 0 {
		return res
	}
	res[len(a)-1] = 1
	tmp := 1
	for i := len(a) - 2; i >= 0; i-- {
		res[i] = res[i+1] * a[i+1]
	}
	for i := 0; i < len(a); i++ {
		res[i] = res[i] * tmp
		tmp *= a[i]
	}
	return res
}

// func constructArr(a []int) []int {
// 	res := make([]int, len(a))
// 	if len(a) == 0 {
// 		return res
// 	}
// 	tmp1 := make([]int, len(a))
// 	tmp2 := make([]int, len(a))
// 	tmp1[len(a)-1] = 1
// 	tmp2[0] = 1
// 	for i := len(a) - 2; i >= 0; i-- {
// 		tmp1[i] = tmp1[i+1] * a[i+1]
// 	}
// 	for i := 1; i < len(a); i++ {
// 		tmp2[i] = tmp2[i-1] * a[i-1]
// 	}
// 	for i := 0; i < len(a); i++ {
// 		res[i] = tmp1[i] * tmp2[i]
// 	}
// 	return res
// }
