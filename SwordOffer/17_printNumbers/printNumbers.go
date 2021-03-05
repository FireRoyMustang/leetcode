package main

import (
	"fmt"
	"strings"
)

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
// 比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func main() {
	fmt.Println(printNumbers(1))
	fmt.Println(printNumbers(2))
	fmt.Println(printNumbers(3))
}

var (
	nums []byte
)

func printNumbers(n int) string {
	res := make([]string, 0)
	nums = make([]byte, 10)
	for i := 0; i < 10; i++ {
		nums[i] = '0' + byte(i)
	}
	dfs(0, n, []byte{}, &res)
	fmt.Println(len(res))
	return strings.Join(res[1:], ",")
}
func dfs(x, n int, s []byte, res *[]string) {
	if x == n-1 {
		for i := 0; i < 10; i++ {
			// if s[0]!='0' 不行，可能越界
			for len(s) > 0 && s[0] == '0' {
				s = s[1:]
			}
			*res = append(*res, string(append(s, nums[i])))
		}
		return
	}
	for i := 0; i < 10; i++ {
		cpy := make([]byte, len(s)+1)
		copy(cpy, s)
		cpy[len(s)] = nums[i]
		dfs(x+1, n, cpy, res)
	}

}

// func printNumbers(n int) []int {
// 	length := 1
// 	for i := 0; i < n; i++ {
// 		length *= 10
// 	}
// 	ans := make([]int, length-1)
// 	for i := 1; i < length; i++ {
// 		ans[i-1] = i
// 	}
// 	return ans
// }
