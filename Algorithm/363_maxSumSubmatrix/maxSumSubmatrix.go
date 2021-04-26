package main

import "fmt"

// 给你一个 m x n 的矩阵 matrix 和一个整数 k ，找出并返回矩阵内部矩形区域的不超过 k 的最大数值和。
// 题目数据保证总会存在一个数值和不超过 k 的矩形区域。
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -100 <= matrix[i][j] <= 100
// -105 <= k <= 105

func main() {
	// fmt.Println(maxSumSubmatrix([][]int{{1, 0, 1}, {0, -2, 3}}, 2))
	// fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 0))
	// fmt.Println(maxSumSubmatrix([][]int{{2, 2, -1}}, 3))
	fmt.Println(maxSumSubmatrix([][]int{{5, -4, -3, 4}, {-3, -4, 4, 5}, {5, 1, 5, -4}}, 8))
}
func maxSumSubmatrix(matrix [][]int, k int) int {
	n, m := len(matrix), len(matrix[0])
	res := -1 << 31
	for l := 0; l < m; l++ { // 左边界
		arr := make([]int, n)
		for r := l; r < m; r++ { // 右边界
			for j := 0; j < n; j++ { // 自上到下
				arr[j] += matrix[j][r]
			}
			res = max(dpMax(arr, k), res)
			if res == k {
				return k
			}
		}
	}
	return res
}
func dpMax(arr []int, k int) int { // 返回数组中连续最大<=k的数
	ans, tmp := -1<<31, 0
	for i := 0; i < len(arr); i++ {
		tmp += arr[i]
		if tmp == k {
			return k
		}
		if tmp > k { // 无效答案
			tmp = 0
		} else {
			ans = max(tmp, ans)
		}
		if tmp < 0 {
			tmp = 0
		}
	}
	fmt.Println(arr, ans)

	return ans
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
