package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "123456"

	fmt.Println(permutation(s))
}

func permutation(s string) []string {
	length := len(s)
	res := make([]string, 0)
	if length == 0 {
		return res
	}
	bytes := []byte(s)
	sort.Sort(byteSlice(bytes))
	nextString := string(bytes)
	for nextString != "" {
		res = append(res, nextString)
		nextString = NextString(bytes)
	}
	return res
}

func NextString(bytes []byte) string {
	length := len(bytes)
	a := length - 2
	for ; a >= 0 && bytes[a] >= bytes[a+1]; a-- {
	}
	if a == -1 {
		return ""
	}
	b := length - 1
	for ; b >= 0 && bytes[b] <= bytes[a]; b-- {
	}
	Swap(bytes, a, b)
	for i, j := a+1, length-1; i < j; i, j = i+1, j-1 {
		Swap(bytes, i, j)
	}
	return string(bytes)
}

//  从小到大排序
type byteSlice []byte

func (bytes byteSlice) Len() int {
	return len(bytes)
}

func (bytes byteSlice) Swap(i, j int) {
	bytes[i], bytes[j] = bytes[j], bytes[i]
}

func (bytes byteSlice) Less(i, j int) bool {
	return bytes[i] < bytes[j]
}

func Swap(bytes []byte, i, j int) {
	if i == j {
		return
	}
	bytes[i], bytes[j] = bytes[j], bytes[i]
}

// func permutation(s string) []string {
// 	length := len(s)
// 	res := make([]string, 0)
// 	if length == 0 {
// 		return res
// 	}
// 	bytes := []byte(s)
// 	dfs(0, bytes, &res)
// 	return res
// }

// func dfs(index int, bytes []byte, res *[]string) {
// 	length := len(bytes)
// 	if index == length-1 {
// 		*res = append(*res, string(bytes))
// 		return
// 	}
// 	set := make(map[byte]bool)
// 	for i := index; i < length; i++ {
// 		if set[bytes[i]] {
// 			continue
// 		}
// 		set[bytes[i]] = true
// 		Swap(bytes, i, index)
// 		dfs(index+1, bytes, res)
// 		Swap(bytes, i, index)
// 	}
// }
