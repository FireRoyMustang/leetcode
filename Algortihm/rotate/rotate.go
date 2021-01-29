package main

/**
原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]

*/
import "fmt"

func main() {
	// matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [4][4]int{{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16}}
	var matrix2 [][]int = make([][]int, 0)
	for i := 0; i < len(matrix); i++ {
		matrix2 = append(matrix2, matrix[i][:])
	}
	rotate(matrix2)
}

// func rotate(matrix [][]int) {
// 	length := len(matrix)
// 	if length == 1 {
// 		return
// 	}
// 	helpMatrix := make([][]int, 0)
// 	var tmp []int
// 	for i := 0; i < length; i++ {
// 		tmp = make([]int, length)
// 		copy(tmp, matrix[i])
// 		helpMatrix = append(helpMatrix, tmp)
// 	}
// 	fmt.Print(helpMatrix)
// 	for i := 0; i < length; i++ {
// 		for j := 0; j < length; j++ {
// 			matrix[j][length-i-1] = helpMatrix[i][j]
// 		}
// 	}
// }
func rotate(matrix [][]int) {
	length := len(matrix)
	if length == 1 {
		return
	}
	if length == 2 {
		tmp := matrix[0][0]
		matrix[0][0] = matrix[1][0]
		matrix[1][0] = matrix[1][1]
		matrix[1][1] = matrix[0][1]
		matrix[0][1] = tmp
	}

	for i := 0; i < length-1; i++ {
		tmpNum := matrix[0][i]
		matrix[0][i] = matrix[length-1-i][0]
		matrix[length-1-i][0] = matrix[length-1][length-1-i]
		matrix[length-1][length-1-i] = matrix[i][length-1]
		matrix[i][length-1] = tmpNum
		fmt.Println(matrix)
		// for j := 0; j < length; j++ {
		// 	fmt.Println(matrix[j])
		// }
	}
	// rotate(matrix[1 : length-1][1 : length-1])
}

// tmp := make([]int, 4)
// tmp[0] = matrix[0][0]
// tmp[1] = matrix[0][length-1]
// tmp[2] = matrix[length-1][0]
// tmp[3] = matrix[length-1][length-1]
