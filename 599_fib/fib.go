package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Printf("fib(%d):%d\n", i, fib(i))
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	pre, cur := 1, 0
	for i := 0; i < n; i++ {
		tmp := cur
		cur += pre
		pre = tmp
	}
	return cur
}
