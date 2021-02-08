package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("numWays(%d):%d\n", i, numWays(i))
	}

}
func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	const mod int = 1e9 + 7
	cur, pre, sum := 1, 0, 0
	for i := 0; i < n; i++ {
		sum = (pre + cur) % mod
		pre, cur = cur, sum
	}
	return sum
}
