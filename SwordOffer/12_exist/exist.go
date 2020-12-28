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
	rows, cols := len(board), len(board[0])
	m := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		m[i] = make([]bool, cols)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backTrace(board, rows, cols, i, j, word, 0, m) {
				return true
			}
		}
	}
	return false
}

// 回溯
func backTrace(board [][]byte, rows, cols, rowStart, colStart int, word string, wordStart int, passed [][]bool) bool {
	if rowStart >= rows || rowStart < 0 || colStart < 0 || colStart >= cols ||
		passed[rowStart][colStart] || word[wordStart] != board[rowStart][colStart] {
		return false
	}
	if wordStart == len(word)-1 {
		return true
	}
	passed[rowStart][colStart] = true
	var flag bool = backTrace(board, rows, cols, rowStart+1, colStart, word, wordStart+1, passed) ||
		backTrace(board, rows, cols, rowStart, colStart+1, word, wordStart+1, passed) ||
		backTrace(board, rows, cols, rowStart-1, colStart, word, wordStart+1, passed) ||
		backTrace(board, rows, cols, rowStart, colStart-1, word, wordStart+1, passed)
	if flag {
		return true
	}
	passed[rowStart][colStart] = false
	return false
}
