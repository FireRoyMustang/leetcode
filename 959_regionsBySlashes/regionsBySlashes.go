package main

// 在由 1 x 1 方格组成的 N x N 网格 grid 中，
// 每个 1 x 1 方块由 /、\ 或空格构成。这些字符会将方块划分为一些共边的区域。
func main() {

}

func regionsBySlashes(grid []string) int {
	N := len(grid)
	res := 0
	if N == 0 {
		return res
	}
	matrix := make([][]bool, N+1)
	for i := 0; i < N+1; i++ {
		matrix[i] = make([]bool, N+1)
	}
	for rowIndex, row := range grid {
		for colIndex := range row {
			switch row[colIndex] {
			case '/':
				if (rowIndex == 0 && colIndex == 0) || (rowIndex == N-1 && colIndex == N-1) {
					res++
				}
				if rowIndex == 0 && matrix[rowIndex+1][colIndex] {
					res++
				}
				if rowIndex
				matrix[rowIndex+1][colIndex] = true
				matrix[rowIndex][colIndex+1] = true
			case '\\':
				if (rowIndex == 0 && colIndex == N-1) || (rowIndex == N-1 && colIndex == 0) {
					res++
				}
				matrix[rowIndex][colIndex] = true
				matrix[rowIndex+1][colIndex+1] = true
			}

		}
	}
	return 1
}
