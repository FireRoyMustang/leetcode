package main

import "fmt"

// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

func main() {
	// fmt.Println(findContinuousSequence(9))
	fmt.Println(findContinuousSequence(15))
}

// 求根法
// func findContinuousSequence(target int) [][]int {
// 	for x := 1; x < target/2; x++ {
// 		delta := 1 - 4*(-x*x+x-2*target)
// 		if delta<0{
// 			continue
// 		}

// 	}
// }
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	left, right, sum := 1, 2, 0
	for left < right {
		sum = (left + right) * (right - left + 1) / 2
		if sum == target {
			tmp := make([]int, right-left+1)
			for i := left; i <= right; i++ {
				tmp[i-left] = i
			}
			res = append(res, tmp)
			left++
		} else if sum < target {
			right++
		} else {
			left++
		}
	}
	return res

}
