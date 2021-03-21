package main

import "fmt"

// 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。
// 请使用原地算法。

// 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个常数空间的解决方案吗？

func main() {
	matrix := [][]int{{0, 1, 2, 0},
		{3, 4, 5, 2},
		{1, 3, 1, 5},
	}
	setZeroes(matrix)
	fmt.Println(matrix)
}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zeroRow := make(map[int]bool)
	zeroCol := make(map[int]bool)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				zeroRow[i] = true
				zeroCol[j] = true
			}
		}
	}
	for row := range zeroRow {
		for i := 0; i < n; i++ {
			matrix[row][i] = 0
		}
	}
	for col := range zeroCol {
		for i := 0; i < m; i++ {
			matrix[i][col] = 0
		}
	}

}
