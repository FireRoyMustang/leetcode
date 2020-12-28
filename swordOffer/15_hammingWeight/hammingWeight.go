package main

import "fmt"

// 请实现一个函数，输入一个整数（以二进制串形式），输出该数二进制表示中 1 的个数。
// 例如，把 9 表示成二进制是 1001，有 2 位是 1。因此，如果输入 9，则该函数输出 2。
func main() {
	i := 9
	fmt.Printf("%d:%d\n", i, hammingWeight(uint32(i)))
}

func hammingWeight(num uint32) int {
	ans := 0
	for num != 0 {
		ans++
		num = num & (num - 1)
	}
	return ans
}
