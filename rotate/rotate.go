package main

func main() {
	matrix := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	var matrix2 [][]int = make([][]int, 0)
	matrix2 = append(matrix2, matrix[0][:])
	matrix2 = append(matrix2, matrix[1][:])
	matrix2 = append(matrix2, matrix[2][:])
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
		// fmt.Println(matrix)
	}
	rotate(matrix[1 : length-1])
}

// tmp := make([]int, 4)
// tmp[0] = matrix[0][0]
// tmp[1] = matrix[0][length-1]
// tmp[2] = matrix[length-1][0]
// tmp[3] = matrix[length-1][length-1]
