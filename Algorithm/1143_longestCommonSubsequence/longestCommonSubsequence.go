package main

import "fmt"

// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。

// 一个字符串的 子序列 是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序
// 的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
// 例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。
// 两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
// 若这两个字符串没有公共子序列，则返回 0。

func main() {
	// fmt.Println(longestCommonSubsequence("abcde", "ace"))
	// fmt.Println(longestCommonSubsequence("abcde", "aec"))
	fmt.Println(longestCommonSubsequence("abcba", "abcbcba"))
}
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	// dp[i][j]表示s1[:i]和s2[:j]的最长公共
	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dp[1][j+1] = dp[0][j] + 1
			} else {
				dp[1][j+1] = max(dp[1][j], dp[0][j+1])
			}
		}
		copy(dp[0], dp[1])
	}
	return dp[1][len(text2)]
}

// func longestCommonSubsequence(text1 string, text2 string) int {
// 	dp := make([][]int, len(text1)+1)
// 	for i := 0; i < len(dp); i++ {
// 		dp[i] = make([]int, len(text2)+1)
// 	}
// 	// dp[i][j]表示s1[:i]和s2[:j]的最长公共
// 	for i := 0; i < len(text1); i++ {
// 		for j := 0; j < len(text2); j++ {
// 			if text1[i] == text2[j] {
// 				dp[i+1][j+1] = dp[i][j] + 1
// 			} else {
// 				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
// 			}
// 		}
// 	}
// 	return dp[len(text1)][len(text2)]
// }
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
