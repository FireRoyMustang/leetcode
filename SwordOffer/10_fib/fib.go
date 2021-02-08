package main

import "fmt"

func main() {
	for i := 1; i < 5; i++ {
		fmt.Printf("fib(%d):%d\n", i, fib(i))
	}
	// i := 95
	// fmt.Printf("fib(%d):%d\n", i, fib(i))

}
func fib(n int) int {
	const mod int = 1e9 + 7
	if n <= 1 {
		return n
	}
	pre := 0
	cur := 1
	for i := 1; i < n; i++ {
		tmp := cur
		cur = (pre + cur) % mod
		pre = tmp
	}
	return cur
}
