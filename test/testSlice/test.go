package main

import "fmt"

var matrix [][]int = make([][]int, 0)

func test(a []int, b int) {
	a = append(a, b)
	if b != 1 {
		matrix = append(matrix, a)
	}
}
func test2(a []int, b int) {
	a = append(a, b)
	fmt.Println("a", a)
}

func main() {
	var s []int = make([]int, 8)
	fmt.Println(cap(s))
	s = append(s, 1)
	fmt.Println(cap(s))
	test2(s, 5)
	test2(s, 10)
	fmt.Println(s)
	// fmt.Println("Matrix:", matrix)

	// fmt.Println(s)
}

// func test(a []int)(b []int){
// 	b=append(a,1,2,3,7)
// 	return
// }
// func main(){
// 	var s []int=[]int{9,10}
// 	s=test(s)
// 	fmt.Println(s)
// }
