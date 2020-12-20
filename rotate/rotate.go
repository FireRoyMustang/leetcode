package main

import "fmt"

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var matrix2 [][]int = make([][]int, 0)
	matrix2 = append(matrix2, matrix[0][:])
	matrix2 = append(matrix2, matrix[1][:])
	matrix2 = append(matrix2, matrix[2][:])
	rotate(matrix2)
}

func rotate(matrix [][]int) {
	length := len(matrix)
	if length == 1 {
		return
	}
	helpMatrix := make([][]int, 0)
	var tmp []int
	for i := 0; i < length; i++ {
		tmp = make([]int, length)
		copy(tmp, matrix[i])
		helpMatrix = append(helpMatrix, tmp)
	}
	fmt.Print(helpMatrix)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			matrix[j][length-i-1] = helpMatrix[i][j]
		}
	}
}
