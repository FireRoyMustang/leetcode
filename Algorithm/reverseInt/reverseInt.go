package main

import "fmt"

func main() {
	fmt.Println(reverse(-10012))
}
func reverse(x int) int {
	const MAXINTDIV10 = (1 << 31) / 10
	ans := 0
	bit := 0
	for {
		if x == 0 {
			break
		}
		bit = x % 10
		if ans > MAXINTDIV10 || (ans == MAXINTDIV10 && bit > 7) {
			return 0
		}
		if ans < -MAXINTDIV10 || (ans == -MAXINTDIV10 && bit < -8) {
			return 0
		}
		ans = ans*10 + bit
		x /= 10
	}
	return ans
}
