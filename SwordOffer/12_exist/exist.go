package main

// 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
// 路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
// 如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
// 例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。

func main() {

}
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	row, col := findChar(word[0], board)
	if row == -1 {
		return false
	}
	for index := 1; index < len(word); index++ {

	}
	return true
}
func findChar(char byte, board [][]byte) (int, int) {
	rows := len(board)
	if rows == 0 {
		return -1, -1
	}
	cols := len(board[0])
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if board[row][col] == char {
				return row, col
			}
		}
	}
	return -1, -1
}

// 回溯
func backTrace(board [][]byte, row, col int, char byte) {

}
