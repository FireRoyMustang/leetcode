package main

import "fmt"

// 给定一个二进制矩阵 A，我们想先水平翻转图像，然后反转图像并返回结果。
// 水平翻转图片就是将图片的每一行都进行翻转，即逆序。例如，水平翻转 [1, 1, 0] 的结果是 [0, 1, 1]。
// 反转图片的意思是图片中的 0 全部被 1 替换， 1 全部被 0 替换。例如，反转 [0, 1, 1] 的结果是 [1, 0, 0]。

func main() {
	fmt.Println(flipAndInvertImage([][]int{
		{1, 1, 0},
		{1, 0, 1},
		{0, 0, 0},
	}))
}

func flipAndInvertImage(A [][]int) [][]int {
	rows, cols := len(A), len(A[0])
	for i := 0; i < rows; i++ {
		left, right := 0, cols-1
		for ; left < right; left, right = left+1, right-1 {
			tmp := A[i][left]
			exchange(A[i][right], &A[i][left])
			exchange(tmp, &A[i][right])
		}
		if left == right {
			exchange(A[i][left], &A[i][right])
		}

	}
	return A
}
func exchange(src int, dst *int) {
	if src == 1 {
		*dst = 0
	} else {
		*dst = 1
	}
}
