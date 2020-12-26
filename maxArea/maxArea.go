package main

func main() {

}
func maxArea(height []int) int {
	length := len(height)
	if length <= 1 {
		return 0
	}
	left := 0
	right := 0
	ans := (height[right] - height[left]) * (right - left + 1)
	for left < right {
		for height[left] >= height[left+1] {
			left++
		}
		if left+1 >= right {
			break
		} else {
			left++
		}
		ans = max(ans, (height[right]-height[left])*(right-left+1))

		for height[right] >= height[right-1] {
			right--
		}
		if left+1 >= right {
			break
		} else {
			right--
		}
		ans = max(ans, (height[right]-height[left])*(right-left+1))
	}
	return ans
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
