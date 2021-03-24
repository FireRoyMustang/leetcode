package main

import (
	"fmt"
)

func main() {
	n := 0          // 员工数
	var w int64 = 0 // 奖金数
	fmt.Scan(&n)
	fmt.Scan(&w)

	ans := 0
	x := make([]int, n) // 奖金下限
	y := make([]int, n) // 奖金上限
	for i := 0; i < n; i++ {
		xi, yi := 0, 0
		fmt.Scan(&xi)
		fmt.Scan(&yi)
		x[i] = xi
		y[i] = yi
	}
	counts := 0
	if n%2 == 0 { // 偶数

	} else { // 奇数

	}
	fmt.Printf("%d\n", ans)
}
