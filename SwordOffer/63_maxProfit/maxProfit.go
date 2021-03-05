package main

import "fmt"

// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？
// 只能买卖一次

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	res := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] <= minPrice {
			minPrice = prices[i]
		} else {
			res = max(prices[i]-minPrice, res)
		}
	}
	return res
}

// func maxProfit(prices []int) int {
// 	if len(prices) == 0 {
// 		return 0
// 	}
// 	res := 0
// 	minPrice := prices[0]
// 	for i := 1; i < len(prices); i++ {
// 		res = max(prices[i]-minPrice, res)
// 		minPrice = min(minPrice, prices[i])
// 	}
// 	return res
// }
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
