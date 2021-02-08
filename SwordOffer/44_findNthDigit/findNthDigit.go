package main

import "fmt"

// 数字以0123456789101112131415…的格式序列化到一个字符序列中。
// 在这个序列中，第5位（从下标0开始计数）是5，第13位是1，第19位是4，等等。
// 请写一个函数，求任意第n位对应的数字。

func main() {
	// fmt.Println(findNthDigit(13))
	// fmt.Println(findNthDigit(19))
	fmt.Println(findNthDigit(20))
}
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	count, res := 10, 0
	pre, digit, bits := count, 10, 2
	for {
		pre = count
		count += (digit*10 - digit) * bits
		if count >= n {
			break
		}
		bits++
		digit *= 10
	}
	diff := n - pre
	res = digit + diff/bits
	res = getKweiNum(res, diff%bits, digit)
	return res
}
func getKweiNum(num, k, digit int) (res int) {
	for i := 0; i <= k; i++ {
		res = num / digit
		num %= digit
		digit /= 10
	}
	return
}
