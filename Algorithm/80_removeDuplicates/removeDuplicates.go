package main

// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 最多出现两次 ，返回删除后数组的新长度。
// 不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。
// 0 <= nums.length <= 3 * 104
// -104 <= nums[i] <= 104
// nums 已按升序排列
func main() {

}

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	prevNum, count, ptr := nums[0], 1, 1
	for i := 1; i < len(nums); i++ {
		if prevNum != nums[i] {
			count = 1
			prevNum = nums[i]
			nums[ptr] = nums[i]
			ptr++
		} else {
			count++
			if count <= 2 {
				nums[ptr] = prevNum
				ptr++
			}
		}
	}
	return ptr
}
