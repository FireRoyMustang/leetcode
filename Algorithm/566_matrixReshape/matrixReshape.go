package main

import "fmt"

// 在MATLAB中，有一个非常有用的函数 reshape，它可以将一个矩阵重塑为另一个大小不同的新矩阵，但保留其原始数据。

// 给出一个由二维数组表示的矩阵，以及两个正整数r和c，分别表示想要的重构的矩阵的行数和列数。

// 重构后的矩阵需要将原始矩阵的所有元素以相同的行遍历顺序填充。

// 如果具有给定参数的reshape操作是可行且合理的，则输出新的重塑矩阵；否则，输出原始矩阵。

func main() {
	nums := [][]int{{1, 2}, {3, 4}}
	fmt.Println(matrixReshape(nums, 1, 4))
}

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rows := len(nums)
	if rows == 0 {
		return nums
	}
	cols := len(nums[0])
	if rows*cols != r*c {
		return nums
	}
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}
	i, j := 0, 0
	for row := 0; row < r; row++ {
		for col := 0; col < c; col++ {
			res[row][col] = nums[i][j]
			j++
			if j == cols {
				j = 0
				i++
			}
		}
	}
	return res
}
