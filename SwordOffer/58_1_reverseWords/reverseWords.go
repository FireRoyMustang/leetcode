package main

import (
	"fmt"
	"strings"
)

// 输入一个英文句子，翻转句子中单词的顺序，但单词内字符的顺序不变。
// 为简单起见，标点符号和普通字母一样处理。例如输入字符串"I am a student. "，则输出"student. a am I"。
func main() {
	// test := "the sky is blue"
	test := "the sky  is  blue"
	// test := "  hello world!  "
	// test := ""
	fmt.Println(reverseWords(test))
}

// func reverseWords(s string) string {
// 	s = strings.Trim(s, " ")
// 	if len(s) == 0 {
// 		return s
// 	}
// 	var build strings.Builder
// 	left, right := len(s)-1, len(s)-1
// 	var lastChar byte = ' '
// 	for i := len(s) - 1; i >= 0; i-- {
// 		if lastChar == ' ' {
// 			if s[i] == ' ' {
// 				continue
// 			} else {
// 				right = i
// 				left = i
// 			}
// 		} else {
// 			if s[i] == ' ' {
// 				build.Write([]byte(s[left : right+1]))
// 				build.WriteByte(s[i])
// 			} else {
// 				left--
// 			}
// 		}
// 		lastChar = s[i]
// 	}
// 	build.Write([]byte(s[left : right+1]))
// 	return build.String()
// }

func reverseWords(s string) string {
	splits := strings.Fields(s)
	if len(splits) == 0 {
		return ""
	}
	var build strings.Builder
	build.WriteString(splits[len(splits)-1])
	for i := len(splits) - 2; i >= 0; i-- {
		build.WriteByte(' ')
		build.WriteString(splits[i])
	}
	return build.String()
}
