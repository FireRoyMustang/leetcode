package main

import "fmt"

func main() {
	testStr := []string{"III", "IV", "IIX", "IX", "LVIII", "MCMXCIV"}
	for _, str := range testStr {
		fmt.Printf("%s: %d\n", str, romanToInt(str))
	}
}
func romanToInt(s string) int {
	prevNum, ans, curNum := 0, 0, 0
	for index := range s {
		curNum = romanCharToInt(s[index])
		if prevNum < curNum {
			ans -= prevNum
		} else {
			ans += prevNum
		}
		prevNum = curNum
	}
	ans += prevNum
	return ans
}
func romanCharToInt(char byte) int {
	switch char {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return -1
}
