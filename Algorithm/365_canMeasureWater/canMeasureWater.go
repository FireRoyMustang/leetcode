package main

import "fmt"

// 有两个容量分别为 x升 和 y升 的水壶以及无限多的水。请判断能否通过使用这两个水壶，
// 从而可以得到恰好 z升 的水？
// 如果可以，最后请用以上水壶中的一或两个来盛放取得的 z升 水。
func main() {
	fmt.Println(canMeasureWater(3, 5, 4)) // true
	fmt.Println(canMeasureWater(2, 6, 5)) // false
}
func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	if x == z || y == z || x+y == z {
		return true
	}

	return z%gcd(x, y) == 0
}

// 求最大公约数
func gcd(x, y int) int {
	remain := x % y
	for remain > 0 {
		x = y
		y = remain
		remain = x % y
	}
	return y
}
