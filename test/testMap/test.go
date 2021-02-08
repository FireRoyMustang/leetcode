package main

import "fmt"

func main() {
	mp := make(map[byte]int)
	fmt.Println(mp[0])
	_, ok := mp[0]
	fmt.Println(ok)
}
