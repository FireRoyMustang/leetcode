package main

import (
	"fmt"
	"math"
)

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Printf("%d:%d\n", i, cuttingRope(i))
	// }
	// fmt.Printf("%d:%d\n", 120, cuttingRope2(120)) // 953271190
	fmt.Printf("%d:%d\n", 127, cuttingRope2(127)) // 439254203

}

// 给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
// 每段绳子的长度记为 k[0],k[1]...k[m-1] 。请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
// 例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，此时得到的最大乘积是18。

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	num, index, base := n%3, float64(n/3), float64(3)
	if num == 0 {
		return int(math.Pow(base, index))
	} else if num == 1 {
		return int(4 * math.Pow(base, index-1))
	}
	return int(2 * math.Pow(base, index))
}

// 答案需要取模 1e9+7（1000000007），如计算初始结果为：1000000008，请返回 1。
const mod int = 1e9 + 7

func cuttingRope2(n int) int {
	if n <= 3 {
		return n - 1
	}
	num, index := n%3, n/3
	if num == 0 {
		return pow(3, index)
	} else if num == 1 {
		return 4 * pow(3, index-1) % mod
	}
	return 2 * pow(3, index) % mod
}
func pow(base, index int) int {
	ans := 1
	// 方法一：依次乘
	// for i := 0; i < index; i++ {
	// 	ans = (ans * 3) % mod
	// }
	// return ans

	// 方法二：二进制取余
	for index > 0 {
		if index%2 != 0 {
			ans = (ans * base) % mod
		}
		base = (base * base) % mod
		index /= 2
	}
	return ans
}
