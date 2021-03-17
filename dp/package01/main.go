package main

// 有N件物品和一个容量为V 的背包。
// 放入第i件物品耗费的空间是Ci，得到的价值是Wi。求解将哪些物品装入背包可使价值总和最大。

func main() {

}
func maxValue(N, V int, C, W []int) int {
	// dp[i][v]代表前i件物品放入容量为v的背包中可以获得的最大价值
	// dp[i][v] = max(dp[i-1][v], dp[i-1][v-ci] + wi)

	// 方法一：未空间优化

	// dp := make([][]int, N+1)
	// for i := 0; i < N+1; i++ {
	// 	dp[i] = make([]int, V+1)
	// }
	// for i := 1; i <= N; i++ {
	// 	ci := C[i-1]
	// 	wi := W[i-1]
	// 	for v := ci; v <= V; v++ {
	// 		dp[i][v] = max(dp[i-1][v], dp[i-1][v-ci]+wi)
	// 	}
	// }
	// return dp[N][V]

	//方法二：空间优化
	// 由于是加前一个数组较小的坐标，因此只要从后向前迭代即可，即保证较小的index后被修改

	dp := make([]int, V+1)
	for i := 0; i < N; i++ {
		for v := V; v >= C[i]; v-- {
			dp[v] = max(dp[v], dp[v-C[i]]+W[i])
		}
	}
	return dp[V]
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
