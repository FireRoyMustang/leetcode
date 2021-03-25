package main

// 数组的每个索引作为一个阶梯，第 i个阶梯对应着一个非负数的体力花费值 cost[i](索引从0开始)。
// 每当你爬上一个阶梯你都要花费对应的体力花费值，然后你可以选择继续爬一个阶梯或者爬两个阶梯。
// 您需要找到达到楼层顶部的最低花费。在开始时，你可以选择从索引为 0 或 1 的元素作为初始阶梯。


// cost 的长度范围是 [2, 1000]。
// cost[i] 将会是一个整型数据，范围为 [0, 999] 。

// dp简化
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	pre, cur := 0, 0
	for i := 2; i <= n; i++ {
		pre, cur = cur, min(cur+cost[i-1], pre+cost[i-2])
	}
	return cur
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
