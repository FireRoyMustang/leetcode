package main

import "fmt"

// 数组传递是值传递，原有数组不会改变

func main() {
	// var arr [5]int // 会初始化为默认值
	// var arr = [5]int{1, 2, 3, 4, 5}
	// var arr = [...]int{1, 2, 3, 4, 5} 这种也是数组
	var arr = [1][1]int{{0}}
	fmt.Println(arr[0][0])
	test(arr)
	fmt.Println(arr[0][0])
}
func test(arr [1][1]int) {
	// 不会改变原有数组的值
	arr[0][0] = 1
}
