package main

import "fmt"

// 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
func main() {
	tests := []string{"(()", ")()())", "", "()(()"}
	for _, test := range tests {
		fmt.Println(longestValidParentheses(test))
	}
}

func longestValidParentheses(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLength = max(maxLength, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
