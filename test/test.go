package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 2, 7, 9, 1}
	sort.Ints(a)
	fmt.Println(a)
}
