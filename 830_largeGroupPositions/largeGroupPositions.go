package main

import "fmt"

// 在一个由小写字母构成的字符串 s 中，包含由一些连续的相同字符所构成的分组。
// 例如，在字符串 s = "abbxxxxzyy" 中，就含有 "a", "bb", "xxxx", "z" 和 "yy" 这样的一些分组。
// 分组可以用区间 [start, end] 表示，其中 start 和 end 分别表示该分组的起始和终止位置的下标。
// 上例中的 "xxxx" 分组用区间表示为 [3,6] 。

// 我们称所有包含大于或等于三个连续字符的分组为 较大分组 。

// 找到每一个 较大分组 的区间，按起始位置下标递增顺序排序后，返回结果。

func main() {
	// testS := "abcdddeeeeaabbbcd"
	testS := "aaa"
	fmt.Println(largeGroupPositions(testS))
}

func largeGroupPositions(s string) [][]int {
	res := make([][]int, 0)
	var preChar byte = ' '
	lens := 0
	for i := 0; i < len(s); i++ {
		if preChar == s[i] {
			lens++
		} else {
			if lens > 2 {
				res = append(res, []int{i - lens, i - 1})
			}
			lens = 1
		}
		preChar = s[i]
	}
	if lens > 2 {
		res = append(res, []int{len(s) - lens, len(s) - 1})
	}
	return res
}
