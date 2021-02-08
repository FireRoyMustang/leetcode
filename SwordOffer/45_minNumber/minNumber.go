package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// 输入一个非负整数数组，把数组里所有数字拼接起来排成一个数，打印能拼接出的所有数字中最小的一个。

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(strconv.FormatInt(int64(i), 2))
	// }
	test := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(test))
}

func minNumber(nums []int) string {
	length := len(nums)
	if length == 1 {
		return strconv.FormatInt(int64(nums[0]), 10)
	}
	strs := make([]string, length)
	var build strings.Builder
	for i := 0; i < length; i++ {
		strs[i] = strconv.FormatInt(int64(nums[i]), 10)
	}
	sort.Sort(strSlice(strs))
	for i := 0; i < length; i++ {
		build.WriteString(strs[i])
	}
	return build.String()
}

type strSlice []string

func (s strSlice) Len() int           { return len(s) }
func (s strSlice) Less(i, j int) bool { return s[i]+s[j] < s[j]+s[i] }
func (s strSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
