package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0

	fmt.Scan(&n)
	arrs := make([][]int, n)
	for i := 0; i < n; i++ {
		m := 0
		fmt.Scan(&m)
		arrs[i] = make([]int, m)
		for j := 0; j < m; j++ {
			mi := 0
			fmt.Scan(&mi)
			arrs[i][j] = mi
		}
		// sort.Ints(arrs[i])
	}
	q := 0
	fmt.Scan(&q)
	res := make([]int, q)

	for i := 0; i < q; i++ {
		n := 0
		fmt.Scan(&n)
		newArr := make([]int, 0)
		index := 0
		for j := 0; j < n; j++ {
			fmt.Scan(&index)
			newArr = append(newArr, arrs[index-1]...)
		}
		// 取n个数组中第k小的数字
		k := 0
		fmt.Scan(&k)
		sort.Ints(newArr)
		res[i] = newArr[k-1]
	}
	for i := 0; i < q; i++ {
		fmt.Printf("%d\n", res[i])
	}
}
