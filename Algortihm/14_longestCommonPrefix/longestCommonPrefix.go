package main

import (
	"fmt"
	"strings"
)

// 编写一个函数来查找字符串数组中的最长公共前缀。
// 如果不存在公共前缀，返回空字符串 ""。

func main() {
	fmt.Println(1)
}
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	var build strings.Builder
	minLen := minLength(strs)
	if length == 0 || minLen == 0 {
		return ""
	}
	if length == 1 {
		return strs[0]
	}
	for i := 0; i < minLen; i++ {
		prevChar := strs[0][i]
		for _, str := range strs {
			if prevChar != str[i] {
				return build.String()
			}
		}
		build.WriteByte(prevChar)
	}
	return build.String()
}

func minLength(strs []string) int {
	min := 1 << 31
	length := 0
	for _, str := range strs {
		length = len(str)
		if min > length {
			min = length
		}
	}
	return min
}
