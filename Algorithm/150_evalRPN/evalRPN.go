package main

import (
	"fmt"
	"strconv"
)

// 根据 逆波兰表示法，求表达式的值。

// 有效的算符包括 +、-、*、/ 。每个运算对象可以是整数，也可以是另一个逆波兰表达式。

// 说明：
// 整数除法只保留整数部分。
// 给定逆波兰表达式总是有效的。换句话说，表达式总会得出有效数值且不存在除数为 0 的情况。
func main() {
	// fmt.Println('+' - '0') // -5
	// fmt.Println('-' - '0') // -3
	// fmt.Println('*' - '0') // -6
	// fmt.Println('/' - '0') // -1
	// fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
	// fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

func evalRPN(tokens []string) int {
	stack := make([]int, 0)
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if token[0] >= '0' || len(token) > 1 { // 数字
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		} else { // 操作符
			length := len(stack)
			num2 := stack[length-1]
			num1 := stack[length-2]
			stack = stack[:length-1]
			switch token[0] {
			case '+':
				stack[length-2] = num1 + num2
			case '-':
				stack[length-2] = num1 - num2
			case '*':
				stack[length-2] = num1 * num2
			case '/':
				stack[length-2] = num1 / num2
			default:
				return 0
			}
		}
	}
	return stack[len(stack)-1]
}
