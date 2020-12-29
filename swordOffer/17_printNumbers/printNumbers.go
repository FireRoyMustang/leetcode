package main

// 输入数字 n，按顺序打印出从 1 到最大的 n 位十进制数。
// 比如输入 3，则打印出 1、2、3 一直到最大的 3 位数 999。
func main() {

}
func printNumbers(n int) []int {
	length := 1
	for i := 0; i < n; i++ {
		length *= 10
	}
	ans := make([]int, length-1)
	for i := 1; i < length; i++ {
		ans[i-1] = i
	}
	return ans
}
