package main

import "fmt"

// 给你一个二维整数数组 matrix， 返回 matrix 的 转置矩阵 。

// 矩阵的 转置 是指将矩阵的主对角线翻转，交换矩阵的行索引与列索引
func main() {
	fmt.Println(transpose([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}

func transpose(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	res := make([][]int, cols)
	for i := 0; i < cols; i++ {
		res[i] = make([]int, rows)
	}
	for row := 0; row < cols; row++ {
		for col := 0; col < rows; col++ {
			res[row][col] = matrix[col][row]
		}
	}
	return res
}
