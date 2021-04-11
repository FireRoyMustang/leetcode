package main

// 给你一个整数 n ，请你判断 n 是否为 丑数 。如果是，返回 true ；否则，返回 false 。
// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。

func main() {

}
func isUgly(n int) bool {
	for n >= 2 {
		if n%2 == 0 {
			n /= 2
		} else {
			break
		}
	}
	for n >= 3 {
		if n%3 == 0 {
			n /= 3
		} else {
			break
		}
	}
	for n >= 5 {
		if n%5 == 0 {
			n /= 5
		} else {
			break
		}
	}
	return n == 1
}
