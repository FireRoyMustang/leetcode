package main

import "fmt"

type Helper struct {
	I int
}

func main() {
	mp := make(map[byte]int)
	fmt.Println(mp[0])
	_, ok := mp[0]
	fmt.Println(ok)

	// map是值传递
	mp2 := make(map[int]Helper)
	mp2[0] = Helper{1}
	helper := mp2[0]
	helper.I = 2
	fmt.Println(mp2[0])
	mp2[0].I = 2
	fmt.Println(mp2[0])

}
