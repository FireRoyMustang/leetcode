package main

import (
	"fmt"
)

func main() {
	N := 0
	fmt.Scan(&N)
	res := make([]byte, N)
	for j := 0; j < N; j++ {
		res[j] = 'Y'
		length := 0
		fmt.Scan(&length)
		arr := make([]int, length)
		for i := 0; i < length; i++ {
			fmt.Scan(&arr[i])
		}
		pre := arr[0]
		flag := false
		preMax := arr[0]
		for i := 1; i < length; i++ {
			if arr[i] > pre {
				if flag { // 超过2次非递增序列
					res[j] = 'N'
					break
				} else {
					flag = true
				}
			}
			pre = arr[i]
		}
		if preMax > arr[length-1] && flag {
			res[j] = 'N'
		}
	}
	for i := 0; i < N; i++ {
		fmt.Printf("%c\n", res[i])
	}
}
