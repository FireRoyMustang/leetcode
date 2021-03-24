package main
// 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。

// 1 <= nums.length <= 3 * 104
// 0 <= nums[i] <= 10^5
func main(){

}

func canJump(nums []int) bool {
	size:=len(nums)
	if size<=1{
		return true
	}
	maxGet:=0
	for i := 0; i < size-1; i++ {
		maxGet=max(maxGet,i+nums[i])
		if maxGet==i{
			return false
		}
		if maxGet>=size-1{
			return true
		}
	}

	return false
}

func max(i,j int)int{
	if i<j{
		return j
	}
	return i
}