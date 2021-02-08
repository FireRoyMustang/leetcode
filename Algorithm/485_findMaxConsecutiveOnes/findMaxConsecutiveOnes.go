package main

// 给定一个二进制数组， 计算其中最大连续1的个数。

func main() {

}
func findMaxConsecutiveOnes(nums []int) int {
	res, tmp := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			res = max(res, tmp)
			tmp = 0
		} else {
			tmp++
		}
	}
	res = max(res, tmp)
	return res
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
