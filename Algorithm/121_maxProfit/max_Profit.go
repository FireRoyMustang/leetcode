package main

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

// 如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

// 注意：你不能在买入股票前卖出股票。

func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	minVal, ans := prices[0], 0
	for i := 1; i < length; i++ {
		if prices[i] > minVal {
			ans = max(ans, prices[i]-minVal)
		} else {
			minVal = prices[i]
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
