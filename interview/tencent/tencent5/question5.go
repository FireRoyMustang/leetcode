package main

import (
	"fmt"
)

func main() {
	T := 0
	fmt.Scan(&T)
	res := make([]int, T)
	for i := 0; i < T; i++ {
		n, m := 0, 0
		fmt.Scan(&n) // 商品件数
		fmt.Scan(&m) // 商品m倍
		w := make([]int, n)
		sum := 0
		for j := 0; j < n; j++ {
			wi := 0
			fmt.Scan(&wi)
			w[i] = wi
			sum += wi
		}
		tmp := sum - sum%m + 1
		dp := make([]bool, tmp+1)
		dp[0] = true
		for j := 0; j < n; j++ {
			for k := w[i]; k <= tmp; k++ {
				if dp[k-w[i]] {
					dp[k] = true
				}
			}
		}
		for j := tmp - 1; j >= 0; j -= m {
			if dp[j] {
				res[i] = sum - j
				break
			}
		}
	}
	for i := 0; i < T; i++ {
		fmt.Printf("%d\n", res[i])
	}
}
