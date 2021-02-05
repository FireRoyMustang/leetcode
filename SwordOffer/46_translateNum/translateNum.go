package main

import (
	"fmt"
	"strconv"
)

// 给定一个数字，我们按照如下规则把它翻译为字符串：0 翻译成 “a” ，1 翻译成 “b”，……，11 翻译成 “l”，
// ……，25 翻译成 “z”。一个数字可能有多个翻译。请编程实现一个函数，用来计算一个数字有多少种不同的翻译方法。

func main() {
	// fmt.Println('a' - '1') 48
	fmt.Println(translateNum(12258))
	// fmt.Println(translateNum(12))
}

func translateNum(num int) int {
	bytes := []byte(strconv.FormatInt(int64(num), 10))
	return recur(bytes)
}

func recur(bytes []byte) int {
	length := len(bytes)
	if length <= 1 {
		return 1
	}
	if bytes[0] == '1' || (bytes[0] == '2' && bytes[1] <= '5') {
		return recur(bytes[2:]) + recur(bytes[1:])
	}
	return recur(bytes[1:])
}
