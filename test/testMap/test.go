package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 4)
	copy(slice2, slice1)
	fmt.Println(slice2)
}
