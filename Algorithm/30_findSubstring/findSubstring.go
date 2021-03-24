package main

import "fmt"

<<<<<<< HEAD
// 给定一个字符串 s 和一些长度相同的单词 words。
// 找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。
// 字符串可能相同

func main() {
	// fmt.Println(findSubstring("barfoothefoobarman", []string{"bar", "foo"}))
	// fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
=======
// 给定一个字符串 s 和一些长度相同的单词 words。找出 s 中恰好可以由 words 中所有单词串联形成的子串的起始位置。
// 注意子串要与 words 中的单词完全匹配，中间不能有其他字符，但不需要考虑 words 中单词串联的顺序。

func main() {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
>>>>>>> 3ce78b6cf5941694602ae8709661090d1be7aea2
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func findSubstring(s string, words []string) []int {
	res := make([]int, 0)
	wordNums, wordLength := len(words), len(words[0])
	wordsMap := make(map[string]int)
	for i := 0; i < len(words); i++ {
<<<<<<< HEAD
		wordsMap[words[i]] = i
	}
	for i := 0; i < len(s)-wordNums*wordLength; i++ {
		recur(s, wordsMap, make([]bool, wordNums), i, 0, wordNums, wordLength, &res)
=======
		wordsMap[words[i]]++
	}
	for i := 0; i <= len(s)-wordNums*wordLength; i++ {
		recur(s, wordsMap, i, 0, wordNums, wordLength, &res)
>>>>>>> 3ce78b6cf5941694602ae8709661090d1be7aea2
	}
	return res
}

<<<<<<< HEAD
func recur(s string, wordsMap map[string]int, mp []bool, start, count, wordNums, wordLength int, res *[]int) {
=======
func recur(s string, wordsMap map[string]int, start, count, wordNums, wordLength int, res *[]int) {
>>>>>>> 3ce78b6cf5941694602ae8709661090d1be7aea2
	if count == wordNums {
		*res = append(*res, start-count*wordLength)
		return
	}
<<<<<<< HEAD
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
=======
	if start+(wordNums-count)*wordLength > len(s) {
		return
	}
	remain, has := wordsMap[s[start:start+wordLength]]
	if !has || remain == 0 {
		return
	}
	wordsMap[s[start:start+wordLength]] = remain - 1
	recur(s, wordsMap, start+wordLength, count+1, wordNums, wordLength, res)
	wordsMap[s[start:start+wordLength]] = remain
>>>>>>> 3ce78b6cf5941694602ae8709661090d1be7aea2
}
