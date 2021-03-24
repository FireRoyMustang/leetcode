package main

import (
	"fmt"
)

func main() {
	T := 0
	fmt.Scan(&T)
	res := make([]int, T)
	for i := 0; i < T; i++ {
		ans := 0
		n := 0
		fmt.Scan(&n)
		for n != 0 {
			if n >= 3 {
				if n%3 == 0 {
					n /= 3
				} else if n%3 == 1 {
					n--
				} else {
					if n%2 == 0 {
						n /= 2
					} else {
						n--
					}
				}
			} else {
				n--
			}
			ans++
		}
		res[i] = ans
	}
	for i := 0; i < T; i++ {
		fmt.Printf("%d\n", res[i])
	}
}
