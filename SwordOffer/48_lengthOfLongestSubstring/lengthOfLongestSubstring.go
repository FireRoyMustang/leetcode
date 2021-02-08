package main

import "fmt"

// 请从字符串中找出一个最长的不包含重复字符的子字符串，计算该最长子字符串的长度。
func main() {
	testStr := []string{"abcabcbb", "bbbbb", "a", "abba", "au"}
	for _, str := range testStr {
		fmt.Println(lengthOfLongestSubstring(str))
	}
	// lengthOfLongestSubstring("abba")
}

// func lengthOfLongestSubstring(s string) int {
// 	length := len(s)
// 	if length <= 1 {
// 		return length
// 	}
// 	mp := make(map[byte]int)
// 	res, dp := 0, 0
// 	for i := 0; i < length; i++ {
// 		index, ok := mp[s[i]]
// 		if !ok {
// 			index = -1
// 		}
// 		if i-index > dp {
// 			dp++
// 		} else {
// 			dp = i - index
// 		}
// 		mp[s[i]] = i
// 		res = max(dp, res)
// 	}
// 	return res
// }

// 双指针
func lengthOfLongestSubstring(s string) int {
	length := len(s)
	mp := make(map[byte]int)
	res, i := 0, -1
	for j := 0; j < length; j++ {
		index, ok := mp[s[j]]
		if ok {
			i = max(index, i)
		}
		mp[s[j]] = j
		res = max(res, j-i)
	}
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
