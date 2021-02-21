package main

// 从扑克牌中随机抽5张牌，判断是不是一个顺子，即这5张牌是不是连续的。
// 2～10为数字本身，A为1，J为11，Q为12，K为13，而大、小王为 0 ，可以看成任意数字。A 不能视为 14。

func main() {

}

// func isStraight(nums []int) bool {
// 	joker := 0
// 	sort.Ints(nums)
// 	for i := 0; i < 4; i++ {
// 		if nums[i] == 0 {
// 			joker++
// 		} else if nums[i] == nums[i+1] {
// 			return false
// 		}
// 	}
// 	return (nums[4] - nums[joker]) < 5
// }
func isStraight(nums []int) bool {
	mi, ma := 14, 0
	set := make([]bool, 13)
	for i := 0; i < 5; i++ {
		if nums[i] == 0 {
			continue
		} else if set[nums[i]] {
			return false
		} else {
			set[nums[i]] = true
			mi = min(nums[i], mi)
			ma = max(nums[i], ma)
		}
	}
	return (ma - mi) < 5
}
func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}
