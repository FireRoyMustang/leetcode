package main

import "fmt"

func main() {
	T := 0
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		res := 0
		N := 0
		fmt.Scan(&N)
		arr := make([]int, N)
		for j := 0; j < N; j++ {
			fmt.Scan(&arr[j])
			res += arr[j]
		}
		u, v := 0, 0
		uv := make([][2]int, N)
		for j := 0; j < N-1; j++ {
			fmt.Scan(&u)
			fmt.Scan(&v)
			uv[j] = [2]int{u, v}
		}
		fmt.Scan("%d\n", res)
	}
}
