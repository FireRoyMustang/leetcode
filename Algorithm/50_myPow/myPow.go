package main

import "fmt"

func main() {
	fmt.Println(myPow(2, 3))
	fmt.Println(myPow(2, -3))
	fmt.Println(myPow(2, 0))

}

func myPow(x float64, n int) float64 {
	res := 1.0
	if n < 0 {
		n = -n
		x = 1 / x
	}
	for n != 0 {
		if n%2 != 0 {
			res *= x
		}
		n >>= 1
		x *= x
	}
	return res
}
