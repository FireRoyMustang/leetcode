package exist

// 请设计一个函数，用来判断在一个矩阵中是否存在一条包含某字符串所有字符的路径。
// 路径可以从矩阵中的任意一格开始，每一步可以在矩阵中向左、右、上、下移动一格。
// 如果一条路径经过了矩阵的某一格，那么该路径不能再次进入该格子。
// 例如，在下面的3×4的矩阵中包含一条字符串“bfce”的路径（路径中的字母用加粗标出）。
var rows int
var cols int
var wordLength int

func exist(board [][]byte, word string) bool {
	wordLength = len(word)
	if wordLength == 0 {
		return true
	}
	rows, cols = len(board), len(board[0])
	bytes := []byte(word)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if backTrace(board, i, j, bytes, 0) {
				return true
			}
		}
	}
	return false
}

// 回溯
func backTrace(board [][]byte, rowStart, colStart int, word []byte, wordStart int) bool {
	if rowStart >= rows || rowStart < 0 || colStart < 0 || colStart >= cols ||
		word[wordStart] != board[rowStart][colStart] {
		return false
	}
	wordStart += 1
	if wordStart == wordLength {
		return true
	}
	board[rowStart][colStart] = 0
	var flag bool = backTrace(board, rowStart+1, colStart, word, wordStart) ||
		backTrace(board, rowStart, colStart+1, word, wordStart) ||
		backTrace(board, rowStart-1, colStart, word, wordStart) ||
		backTrace(board, rowStart, colStart-1, word, wordStart)
	if flag {
		return true
	}
	board[rowStart][colStart] = word[wordStart-1]
	return false
}
