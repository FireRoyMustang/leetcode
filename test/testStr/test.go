package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a  b  c"
	// 使用Fields()
	ss := strings.Split(s, " ")
	for i := 0; i < len(ss); i++ {
		fmt.Printf("Splits[%d]:%v|\n", i, ss[i])
	}
}
