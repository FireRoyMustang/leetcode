package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// 请你来实现一个 atoi 函数，使其能将字符串转换成整数。

// 首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。接下来的转化规则如下：

// 如果第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字字符组合起来，
// 形成一个有符号整数。
// 假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成一个整数。
// 该字符串在有效的整数部分之后也可能会存在多余的字符，那么这些字符可以被忽略，
// 它们对函数不应该造成影响。

func main() {
	testStr := "4193 with words"
	fmt.Println(myAtoi(testStr))
}
func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	s = strings.TrimFunc(s, unicode.IsLetter)
	fmt.Println(s)
	ans, _ := strconv.Atoi(s)
	return ans
}
