package main

// 实现函数double Power(double base, int exponent)，
// 求base的exponent次方。不得使用库函数，同时不需要考虑大数问题。

import "fmt"

func main() {
	i := -2
	fmt.Printf("%v:%v\n", i, myPow(2.0, i))
}
func myPow(x float64, n int) float64 {
	if x == 0 {
		return 0
	}
	var ans float64 = 1
	if n < 0 {
		x = 1 / x
		n = -n
	}
	for n != 0 {
		if n%2 != 0 {
			ans *= x
		}
		x *= x
		n >>= 1
	}
	return ans
}
