package main

import "fmt"

// 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。
// 你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。
// 给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

func main() {
	test := make([][]int, 3)
	test[0] = []int{1, 3, 1}
	test[1] = []int{1, 5, 1}
	test[2] = []int{4, 2, 1}
	fmt.Println(maxValue(test))
}

// func maxValue(grid [][]int) int {
// 	rows := len(grid)
// 	if rows == 0 {
// 		return 0
// 	}
// 	cols := len(grid[0])
// 	if cols == 0 {
// 		return 0
// 	}
// 	dp := make([][]int, rows+1)
// 	for i := 0; i <= rows; i++ {
// 		dp[i] = make([]int, cols+1)
// 	}
// 	for row := 1; row <= rows; row++ {
// 		for col := 1; col <= cols; col++ {
// 			dp[row][col] = max(dp[row-1][col]+grid[row-1][col-1], dp[row][col-1]+grid[row-1][col-1])
// 		}
// 	}
// 	return dp[rows][cols]
// }

func maxValue(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	if cols == 0 {
		return 0
	}
	dp := make([]int, cols+1)
	for row := 1; row <= rows; row++ {
		for col := 1; col <= cols; col++ {
			dp[col] = max(dp[col]+grid[row-1][col-1], dp[col-1]+grid[row-1][col-1])
		}
	}
	return dp[cols]
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
