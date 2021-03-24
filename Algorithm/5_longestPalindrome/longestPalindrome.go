package main

// 给你一个字符串 s，找到 s 中最长的回文子串。

import "fmt"

func main() {
	// longestPalindrome("babad")
	fmt.Println(longestPalindrome("ac"))
}
func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	dp := make([][]bool, length)
	for i := 0; i < length; i++ {
		dp[i] = make([]bool, length)
	}
	maxLen := -1
	ans := ""
	for l := 0; l < length; l++ { // 代表回文串的长度
		for i := 0; i < length-l; i++ {
			j := i + l
			if l == 0 {
				dp[i][j] = true
			} else if l == 1 {
				dp[i][j] = s[i] == s[j]
			} else {
				dp[i][j] = dp[i+1][j-1] && (s[i] == s[j])
			}
			if l > maxLen && dp[i][j] {
				maxLen = l
				ans = string(s[i : j+1])
			}
		}
	}
	return ans
}
