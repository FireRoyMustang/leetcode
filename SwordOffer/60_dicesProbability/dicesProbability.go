package main

import "fmt"

// 把n个骰子扔在地上，所有骰子朝上一面的点数之和为s。输入n，打印出s的所有可能的值出现的概率。
// 你需要用一个浮点数数组返回答案，其中第 i 个元素代表这 n 个骰子所能掷出的点数集合中第 i 小的那个的概率。

func main() {
	fmt.Println(dicesProbability(2))
}

// 方法一
func dicesProbability(n int) []float64 {
	length := 5*n + 1
	res := make([]float64, length)
	if n == 1 {
		for i := 0; i < length; i++ {
			res[i] = 1.0 / 6
		}
		return res
	}
	lastProb := dicesProbability(n - 1)
	for i := 0; i < len(lastProb); i++ {
		for j := 0; j < 6; j++ { // 点数
			res[i+j] += lastProb[i] / 6
		}
	}
	return res
}
