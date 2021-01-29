package main

import (
	"fmt"
)

// 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(generateParenthesis(i))
	}
}

func generateParenthesis(n int) []string {
	var res []string = make([]string, 0)
	// left和right分别代表可以加的左括号与右括号
	left, right := n, 0
	helper(&res, "", left, right)
	return res
}

func helper(res *[]string, s string, left, right int) {
	if left == 0 && right == 0 {
		*res = append(*res, s)
		return
	}
	// var builds1, builds2 strings.Builder
	// if left > 0 {
	// 	builds1.WriteString(s)
	// 	builds1.WriteByte('(')
	// 	helper(res, builds1.String(), left-1, right+1)
	// }
	// if right > 0 {
	// 	builds2.WriteString(s)
	// 	builds2.WriteByte(')')
	// 	helper(res, builds2.String(), left, right-1)
	// }
	if left > 0 {
		helper(res, s+"(", left-1, right+1)
	}
	if right > 0 {
		helper(res, s+")", left, right-1)
	}
}
