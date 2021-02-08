package main

import "fmt"

// 给你两个长度相同的字符串，s 和 t。
// 将 s 中的第 i 个字符变到 t 中的第 i 个字符需要 |s[i] - t[i]| 的开销（开销可能为 0），也就是两个字符的 ASCII 码值的差的绝对值。
// 用于变更字符串的最大预算是 maxCost。在转化字符串时，总开销应当小于等于该预算，这也意味着字符串的转化可能是不完全的。
// 如果你可以将 s 的子字符串转化为它在 t 中对应的子字符串，则返回可以转化的最大长度。
// 如果 s 中没有子字符串可以转化成 t 中对应的子字符串，则返回 0。

func main() {
	// s, t, cost := "abcd", "bcdf", 3
	s, t, cost := "krrgw", "zjxss", 19
	// s, t, cost := "abcd", "cdef", 3
	fmt.Println(equalSubstring(s, t, cost))
}

func equalSubstring(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] > t[i] {
			diff[i] = int(s[i] - t[i])
		} else {
			diff[i] = int(t[i] - s[i])
		}
	}
	sum, start := 0, 0
	for end, d := range diff {
		sum += d
		for sum > maxCost {
			sum -= diff[start]
			start++
		}
		maxLen = max(maxLen, end-start+1)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func equalSubstring(s string, t string, maxCost int) int {
// 	length := len(s)
// 	mp := make([]int, 26)
// 	res := 0
// 	for i := 0; i < length; i++ {
// 		if s[i] > t[i] {
// 			mp[s[i]-t[i]]++
// 		} else {
// 			mp[t[i]-s[i]]++
// 		}
// 	}
// 	for i := 0; i < 26; i++ {
// 		fmt.Printf("i:%d count:%d\n", i, mp[i])
// 		cost := mp[i] * i
// 		if maxCost >= cost {
// 			res += mp[i]
// 			maxCost -= cost
// 		} else {
// 			res += maxCost / i
// 			break
// 		}
// 	}
// 	return res
// }
