package main

/**
给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。
你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

*/
import "fmt"

func main() {
	// matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}

	rotate(matrix)
	fmt.Println(matrix)
}

// func rotate(matrix [][]int) {
// 	side := len(matrix)
// 	if side == 1 {
// 		return
// 	}
// 	helpMatrix := make([][]int, 0)
// 	var tmp []int
// 	for i := 0; i < side; i++ {
// 		tmp = make([]int, side)
// 		copy(tmp, matrix[i])
// 		helpMatrix = append(helpMatrix, tmp)
// 	}
// 	fmt.Print(helpMatrix)
// 	for i := 0; i < side; i++ {
// 		for j := 0; j < side; j++ {
// 			matrix[j][side-i-1] = helpMatrix[i][j]
// 		}
// 	}
// }

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}
