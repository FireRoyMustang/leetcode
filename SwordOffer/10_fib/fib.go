package main

import "fmt"

func main() {
	// for i := 20; i < 46; i++ {
	// 	fmt.Printf("fib(%d):%d\n", i, fib(i))
	// }
	i := 95
	fmt.Printf("fib(%d):%d\n", i, fib(i))

}
func fib(n int) int {
	const mod int = 1e9 + 7
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	pre := 0
	cur := 1
	for i := 0; i < n-1; i++ {
		tmp := cur
		cur = (pre + cur) % mod
		pre = tmp
	}
	return cur
}
