package main

import (
	"fmt"
	"strings"
)

// n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
// 给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案。
// 每一种解法包含一个不同的 n 皇后问题 的棋子放置方案，该方案中 'Q' 和 '.' 分别代表了皇后和空位。

func main() {
	fmt.Println(solveNQueens(4))
}

var (
	queens    []int
	columns   []bool
	leftDiag  map[int]bool
	rightDiag map[int]bool
	ns        int
)

func solveNQueens(n int) [][]string {
	ns = n
	res := make([][]string, 0)
	columns = make([]bool, n)
	leftDiag = map[int]bool{}
	rightDiag = map[int]bool{}
	queens = make([]int, n)
	backtrace(0, &res)
	return res

}
func backtrace(row int, res *[][]string) {
	if row == ns {
		*res = append(*res, generateBoard())
		return
	}
	for i := 0; i < ns; i++ {
		if judgeCollide(row, i) {
			left := row - i
			right := row + i
			columns[i] = true
			leftDiag[left] = true
			rightDiag[right] = true

			queens[row] = i
			backtrace(row+1, res)
			queens[row] = -1

			columns[i] = false
			leftDiag[left] = false
			rightDiag[right] = false
		}

	}

}

func judgeCollide(row, col int) bool {
	// 判断行是否合理，没有必要
	// 判断列
	if columns[col] {
		return false
	}
	// 判断左斜
	if leftDiag[row-col] {
		return false
	}
	// 判断右斜
	if rightDiag[row+col] {
		return false
	}
	return true
}

func generateBoard() []string {
	res := make([]string, ns)
	var build strings.Builder
	for row := 0; row < ns; row++ {
		col := queens[row]
		for i := 0; i < col; i++ {
			build.WriteByte('.')
		}
		build.WriteByte('Q')
		for i := col + 1; i < ns; i++ {
			build.WriteByte('.')
		}

		res[row] = build.String()
		build.Reset()
	}

	return res
}

// func generateBoard(col int) string {
// 	var build strings.Builder
// 	for i := 0; i < col; i++ {
// 		build.WriteByte('.')
// 	}
// 	build.WriteByte('Q')
// 	for i := col + 1; i < ns; i++ {
// 		build.WriteByte('.')
// 	}
// 	return build.String()
// }

/* var solutions [][]string

func solveNQueens(n int) [][]string {
	solutions = [][]string{}
	queens := make([]int, n)
	for i := 0; i < n; i++ {
		queens[i] = -1
	}
	columns := map[int]bool{}
	diagonals1, diagonals2 := map[int]bool{}, map[int]bool{}
	backtrack(queens, n, 0, columns, diagonals1, diagonals2)
	return solutions
}

func backtrack(queens []int, n, row int, columns, diagonals1, diagonals2 map[int]bool) {
	if row == n {
		board := generateBoard(queens, n)
		solutions = append(solutions, board)
		return
	}
	for i := 0; i < n; i++ {
		if columns[i] {
			continue
		}
		diagonal1 := row - i
		if diagonals1[diagonal1] {
			continue
		}
		diagonal2 := row + i
		if diagonals2[diagonal2] {
			continue
		}
		queens[row] = i
		columns[i] = true
		diagonals1[diagonal1], diagonals2[diagonal2] = true, true
		backtrack(queens, n, row+1, columns, diagonals1, diagonals2)
		queens[row] = -1
		delete(columns, i)
		delete(diagonals1, diagonal1)
		delete(diagonals2, diagonal2)
	}
}

func generateBoard(queens []int, n int) []string {
	board := []string{}
	for i := 0; i < n; i++ {
		row := make([]byte, n)
		for j := 0; j < n; j++ {
			row[j] = '.'
		}
		row[queens[i]] = 'Q'
		board = append(board, string(row))
	}
	return board
} */
