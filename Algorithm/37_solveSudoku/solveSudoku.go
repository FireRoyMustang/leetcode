package main

import "fmt"

// 编写一个程序，通过填充空格来解决数独问题。

// 一个数独的解法需遵循如下规则：
// 数字 1-9 在每一行只能出现一次。
// 数字 1-9 在每一列只能出现一次。
// 数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
// 空白格用 '.' 表示

// 给定的数独序列只包含数字 1-9 和字符 '.' 。
// 你可以假设给定的数独只有唯一解。
// 给定数独永远是 9x9 形式的。

func main() {
	a := make([]int, 10)
	a['9'-'1'] = 1
	fmt.Println(a[8])
}

func solveSudoku(board [][]byte) {
	var rows, cols [9][9]bool
	var blocks [3][3][9]bool
	spaces := make([][2]int, 0)
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] != '.' {
				num := board[row][col] - '1'
				rows[row][num] = true
				cols[col][num] = true
				blocks[row/3][col/3][num] = true
			} else {
				spaces = append(spaces, [2]int{row, col})
			}
		}
	}
	backTrace(board, 0, rows, cols, blocks, spaces)
}

func backTrace(board [][]byte, ptr int, rows, cols [9][9]bool, blocks [3][3][9]bool, spaces [][2]int) bool {
	if ptr == len(spaces) {
		return true
	}
	row, col := spaces[ptr][0], spaces[ptr][1]
	// 填数字
	for num := byte(0); num < 9; num++ {
		// 判断合法
		if rows[row][num] || cols[col][num] || blocks[row/3][col/3][num] {
			continue
		}
		board[row][col] = '1' + num
		rows[row][num] = true
		cols[col][num] = true
		blocks[row/3][col/3][num] = true
		if backTrace(board, ptr+1, rows, cols, blocks, spaces) {
			return true
		} else {
			rows[row][num] = false
			cols[col][num] = false
			blocks[row/3][col/3][num] = false
		}
	}
	return false
}
