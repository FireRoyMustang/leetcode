package main

// 给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。
// 请使用原地算法。

// 一个直接的解决方案是使用  O(mn) 的额外空间，但这并不是一个好的解决方案。
// 一个简单的改进方案是使用 O(m + n) 的额外空间，但这仍然不是最好的解决方案。
// 你能想出一个常数空间的解决方案吗？

func main() {

}

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	zerosRow := make(map[int]bool)
	zerosCol := make(map[int]bool)
	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {

			}
		}
	}
}
