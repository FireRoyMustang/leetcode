package main

import (
	"fmt"
)

// 森林中，每个兔子都有颜色。其中一些兔子（可能是全部）告诉你还有多少其他的兔子和自己有相同的颜色。
// 我们将这些回答放在 answers 数组里。
// 返回森林中兔子的最少数量。
// answers 的长度最大为1000。
// answers[i] 是在 [0, 999] 范围内的整数。
func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
	fmt.Println(numRabbits([]int{10, 10, 10}))
	fmt.Println(numRabbits([]int{1, 0, 1, 0, 0}))
}

// func numRabbits(answers []int) int {
// 	size := len(answers)
// 	if size == 0 {
// 		return 0
// 	}
// 	res := 0
// 	countMap := make(map[int]int)
// 	for i := 0; i < size; i++ {
// 		count, ok := countMap[answers[i]]
// 		if ok {
// 			countMap[answers[i]] = count + 1
// 		} else {
// 			countMap[answers[i]] = 1
// 		}
// 	}
// 	for k, v := range countMap {
// 		if v%(k+1) == 0 {
// 			res += v
// 		} else {
// 			res += v/(k+1)*(k+1) + k + 1
// 		}
// 	}
// 	return res
// }
func numRabbits(answers []int) int {
	size := len(answers)
	if size == 0 {
		return 0
	}
	res := 0
	countMap := make(map[int]uint16)
	for i := 0; i < size; i++ {
		count, ok := countMap[answers[i]]
		if ok {
			countMap[answers[i]] = count + 1
		} else {
			countMap[answers[i]] = 1
		}
	}
	for k, b := range countMap {
		v := int(b)
		if v%(k+1) == 0 {
			res += v
		} else {
			res += v/(k+1)*(k+1) + k + 1
		}
	}
	return res
}
