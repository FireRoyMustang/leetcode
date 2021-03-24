package main

import "fmt"

// 给定一个字符串 s 和一些长度相同的单词 words。
// 找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
// 字符串可能相同

func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	wordNums, wordLength := len(words), len(words[0])
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordsMap[words[i]] = i
	}
	for i := 0; i < len(s)-wordNums*wordLength; i++ {
		recur(s, wordsMap, make([]bool, wordNums), i, 0, wordNums, wordLength, &res)
	}
	return res
}

func recur(s string, wordsMap map[string]int, mp []bool, start, count, wordNums, wordLength int, res *[]int) {
	if count == wordNums {
		*res = append(*res, start-count*wordLength)
		return
	}
	if start+wordLength > len(s) {
		return
	}
	index, has := wordsMap[s[start:start+wordLength]]
	if !has {
		return
	}
	if mp[index] {
		return
	}
	mp[index] = true
	recur(s, wordsMap, mp, start+wordLength, count+1, wordNums, wordLength, res)
	mp[index] = false
}
