package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 7, 9, 1}
	sort.Ints(a)
	fmt.Println(a)
	fmt.Printf("%d\t%d\t%d\n", '1', 'z'-'0', 'A')
	s := "abcdefg"
	s[0] = 'c'
	fmt.Println(s)
}
