package main

import "fmt"

// 给定一个字符串 s 和一个字符串 t ，计算在 s 的子序列中 t 出现的个数。
// 字符串的一个 子序列 是指，通过删除一些（也可以不删除）字符且不干扰剩余字符
// 相对位置所组成的新字符串。（例如，"ACE" 是 "ABCDE" 的一个子序列，而 "AEC" 不是）
// 题目数据保证答案符合 32 位带符号整数范围。

// 0 <= s.length, t.length <= 1000
// s 和 t 由英文字母组成
func main() {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
	fmt.Println(numDistinct("babgbag", "bag"))
}

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i <= len(t); i++ {
		dp[i] = make([]int, len(s)+1)
	}
	// dp[i][j]表示t[:i]与s[:j]匹配的个数
	for i := 0; i <= len(s); i++ {
		dp[0][i] = 1
	}
	for i := 1; i <= len(t); i++ {
		for j := 1; j <= len(s); j++ {
			if t[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[len(t)][len(s)]
}

// func numDistinct(s string, t string) int {
// 	res := 0
// 	bytes := []byte(s)
// 	res += dfs(bytes, t, 0)
// 	return res
// }

// func dfs(bytes []byte, s string, start int) int {
// 	if start == len(s) {
// 		return 1
// 	}
// 	res := 0
// 	if len(bytes) == 0 {
// 		return res
// 	}
// 	for i := 0; i < len(bytes)-(len(s)-start)+1; i++ {
// 		if bytes[i] == s[start] {
// 			res += dfs(bytes[i+1:], s, start+1)
// 		}
// 	}
// 	return res
// }
