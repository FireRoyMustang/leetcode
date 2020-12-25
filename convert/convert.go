package main

// 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
// 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
func main() {
	convert("AB", 1)
}
func convert(s string, numRows int) string {
	length := len(s)
	if numRows == 1 {
		return s
	}
	ans := make([][]byte, min(numRows, length))
	for i := 0; i < len(ans); i++ {
		ans[i] = make([]byte, 0)
	}
	curRow := 0
	direction := true
	for i := 0; i < length; i++ {
		ans[curRow] = append(ans[curRow], s[i])
		if direction {
			curRow++
			if curRow == numRows-1 {
				direction = false
			}
		} else {
			curRow--
			if curRow == 0 {
				direction = true
			}
		}

	}
	result := ""
	for _, bytes := range ans {
		result += string(bytes)
	}
	return result
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
