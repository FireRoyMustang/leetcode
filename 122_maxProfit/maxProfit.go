package main

import "fmt"

// 给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

// 设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。
func main() {
	test := []int{1, 2, 3, 4, 5}
	fmt.Printf("test %v:%d\n", test, maxProfit(test))
}
func maxProfit(prices []int) int {
	length := len(prices)
	if length <= 1 {
		return 0
	}
	const MAXINT int = 1 << 31
	minNum, ans := prices[0], 0
	for i := 1; i < length; i++ {
		if prices[i] > minNum {
			ans += prices[i] - minNum
			minNum = prices[i]
		} else {
			minNum = prices[i]
		}
	}
	return ans
}
