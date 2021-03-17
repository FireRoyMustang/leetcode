package main

import "fmt"

// 集团里有 n 名员工，他们可以完成各种各样的工作创造利润。
// 第 i 种工作会产生 profit[i] 的利润，它要求 group[i] 名成员共同参与。
// 如果成员参与了其中一项工作，就不能参与另一项工作。
// 工作的任何至少产生 minProfit 利润的子集称为盈利计划。并且工作的成员总数最多为 n 。
// 有多少种计划可以选择？因为答案很大，所以 返回结果模 10^9 + 7 的值。

func main() {
	fmt.Println(profitableSchemes(5, 3, []int{2, 2}, []int{2, 3}))
}

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	const MOD int = 1e9 + 7
	size := len(group)
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}
	for i := 1; i <= size; i++ {
		// dp[j][k]表示到i序号为止，j名罪犯获取至少k的利益的方式数
		g := group[i-1]
		p := profit[i-1]
		for j := n; j >= g; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] += dp[j-g][max(k-p, 0)]
				dp[j][k] %= MOD
			}
		}
	}
	// fmt.Println(dp)
	return dp[n][minProfit]
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
