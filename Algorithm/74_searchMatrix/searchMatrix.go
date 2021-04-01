package main

// 编写一个高效的算法来判断 m x n 矩阵中，是否存在一个目标值。该矩阵具有如下特性：

// 每行中的整数从左到右按升序排列。
// 每行的第一个整数大于前一行的最后一个整数。
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10^4 <= matrix[i][j], target <= 10^4

func main() {

}
func searchMatrix(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	left, right := 0, m*n
	for left < right {
		mid := (left + right) >> 1
		row, col := calCoordinate(mid, m, n)
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			right = mid
		} else { //<
			left = mid + 1
		}
	}
	return false
}
func calCoordinate(index, m, n int) (row, col int) {
	row = index / n
	col = index - row*n
	return
}
