package main

import "fmt"

func main() {
	mp1 := make(map[int]int)
	mp1[1]++
	fmt.Println(mp1[1])
}
