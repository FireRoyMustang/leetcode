package main

// 给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。
// 你可以假定该字符串只包含小写字母
import "fmt"

func main() {
	fmt.Println(firstUniqChar("helloworld"))
}
func firstUniqChar(s string) int {
	length := len(s)
	if length == 0 {
		return -1
	}
	if length == 1 {
		return 0
	}
	m := make(map[byte]bool, 0)
	arr := make([]int, 26)
	for i := 0; i < length; i++ {
		isFirst, appear := m[s[i]]
		if !appear {
			m[s[i]] = true
			arr[s[i]-'a'] = i
		} else if isFirst {
			m[s[i]] = false
		}
	}
	ans := -1
	for char, isFirst := range m {
		if isFirst {
			if ans == -1 {
				ans = arr[char-'a']
			} else {
				ans = min(ans, arr[char-'a'])
			}
		}
	}
	return ans
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func firstUniqChar2(s string) int {
	cnt := [26]int{}
	for _, ch := range s {
		cnt[ch-'a']++
	}
	for i, ch := range s {
		if cnt[ch-'a'] == 1 {
			return i
		}
	}
	return -1
}
