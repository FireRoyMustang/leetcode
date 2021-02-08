package main

// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

func main() {

}

func firstUniqChar(s string) byte {
	length := len(s)
	set := make([]int, 26)
	if length == 1 {
		return s[0]
	}
	for i := 0; i < length; i++ {
		set[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if set[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
