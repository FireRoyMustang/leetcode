package main

import "fmt"

// 给你一个 m x n 的矩阵 matrix 。如果这个矩阵是托普利茨矩阵，返回 true ；否则，返回 false 。
// 如果矩阵上每一条由左上到右下的对角线上的元素都相同，那么这个矩阵是 托普利茨矩阵 。

// 进阶问题
// 如果矩阵存储在磁盘上，并且内存有限，以至于一次最多只能将矩阵的一行加载到内存中，该怎么办？
// 如果矩阵太大，以至于一次只能将不完整的一行加载到内存中，该怎么办？

// 对于进阶问题一，一次最多只能将矩阵的一行加载到内存中，我们将每一行复制到一个连续数组中，
// 随后在读取下一行时，就与内存中此前保存的数组进行比较。

// 对于进阶问题二，一次只能将不完整的一行加载到内存中，我们将整个矩阵竖直切分成若干子矩阵，
// 并保证两个相邻的矩阵至少有一列或一行是重合的，然后判断每个子矩阵是否符合要求。

func main() {
	fmt.Println(isToeplitzMatrix([][]int{
		{36, 59, 71, 15, 26, 82, 87},
		{56, 36, 59, 71, 15, 26, 82},
		{15, 0, 36, 59, 71, 15, 26},
	}))
}

func isToeplitzMatrix(matrix [][]int) bool {
	rows := len(matrix)
	if rows <= 1 {
		return true
	}
	cols := len(matrix[0])
	if cols <= 1 {
		return true
	}
	for col := 1; col < cols; col++ {
		for row := 1; row < rows; row++ {
			if matrix[row][col] != matrix[row-1][col-1] {
				return false
			}
		}
	}
	return true
}
