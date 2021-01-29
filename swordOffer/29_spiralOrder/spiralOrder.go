package main

// 输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。

func main() {

}

func spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return make([]int, 0)
	}
	cols := len(matrix[0])
	res := make([]int, rows*cols)
	if cols == 0 {
		return make([]int, 0)
	}
	up, down, left, right, ptr := 0, rows-1, 0, cols-1, 0
	for {
		for i := left; i <= right; i++ {
			res[ptr] = matrix[up][i]
			ptr++
		}
		up++
		if up > down {
			break
		}
		for i := up; i <= down; i++ {
			res[ptr] = matrix[i][right]
			ptr++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			res[ptr] = matrix[down][i]
			ptr++
		}
		down--
		if up > down {
			break
		}
		for i := down; i >= up; i-- {
			res[ptr] = matrix[i][left]
			ptr++
		}
		left++
		if left > right {
			break
		}
	}
	return res
}
