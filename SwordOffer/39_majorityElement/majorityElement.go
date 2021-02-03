package main

// 数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。

func main() {

}
func majorityElement(nums []int) int {
	votes, maj := 0, 0
	for _, num := range nums {
		if votes == 0 {
			maj = num
			votes++
		} else {
			if num != maj {
				votes--
			} else {
				votes++
			}
		}
	}
	return maj
}
