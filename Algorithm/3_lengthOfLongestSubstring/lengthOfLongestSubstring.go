package main

import "fmt"

//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
// 0 <= s.length <= 5 * 104
// s 由英文字母、数字、符号和空格组成
func main() {
	// s := "abcabcbb"
	s := "bbbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
func lengthOfLongestSubstring(s string) int {
	// 表明字符出现的最后位置
	indexMap := make(map[byte]int)
	res, start := 0, -1
	for i := 0; i < len(s); i++ {
		index, has := indexMap[s[i]]
		if has && index > start {
			start = index
		} else {
			res = max(res, i-start)
		}
		indexMap[s[i]] = i
	}
	return res
}

// func lengthOfLongestSubstring(s string) int {
// 	// 表明字符出现的最后位置
// 	indexMap := make(map[byte]int)
// 	res, start := 0, -1
// 	for i := 0; i < len(s); i++ {
// 		index, has := indexMap[s[i]]
// 		if has {
// 			start = max(start, index)
// 		}
// 		indexMap[s[i]] = i
// 		res = max(res, i-start)
// 	}
// 	return res
// }

// func lengthOfLongestSubstring(s string) int {
// 	map1 := map[uint8]int{}
// 	length := len(s)
// 	if length <= 1 {
// 		return length
// 	}
// 	start := 0
// 	end := 0
// 	result := 0
// 	for i := 0; i < length-1; i++ {
// 		id, contains := map1[s[i]]
// 		if contains {
// 			result = max(result, end-start)
// 			deleteMap(s, map1, start, id)
// 			start = id + 1
// 		}
// 		end++
// 		map1[s[i]] = i
// 	}
// 	_, contains := map1[s[length-1]]
// 	if contains {
// 		result = max(result, end-start)
// 	} else {
// 		result = max(result, end-start+1)
// 	}
// 	end++

// 	return result

// }

// func lengthOfLongestSubstring1(s string) int {
// 	// 哈希集合，记录每个字符是否出现过
// 	m := map[byte]int{}
// 	n := len(s)
// 	// 右指针，初始值为 -1，相当于我们在字符串的左边界的左侧，还没有开始移动
// 	rk, ans := -1, 0
// 	for i := 0; i < n; i++ {
// 		if i != 0 {
// 			// 左指针向右移动一格，移除一个字符
// 			delete(m, s[i-1])
// 		}
// 		for rk+1 < n && m[s[rk+1]] == 0 {
// 			// 不断地移动右指针
// 			m[s[rk+1]]++
// 			rk++
// 		}
// 		// 第 i 到 rk 个字符是一个极长的无重复字符子串
// 		ans = max(ans, rk-i+1)
// 	}
// 	return ans
// }
// func deleteMap(s string, map1 map[uint8]int, start, end int) {
// 	for i := start; i <= end; i++ {
// 		delete(map1, s[i])
// 	}
// }
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
