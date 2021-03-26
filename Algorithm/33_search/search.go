package main 

// 整数数组 nums 按升序排列，数组中的值 互不相同 。

// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。

// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的索引，否则返回 -1 。


func main() {
	search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
}

func search(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) >>1
		if nums[mid] == target{
			return mid
		}
		// 判断是否有序
		if nums[mid] >= nums[0]{ // 左边有序
			if target < nums[mid] && target >=nums[left]{
				right = mid
			}else {
				left = mid + 1
			}
		}else{ // 右边有序
			if target > nums[mid] && target <=nums[right-1]{
				left = mid + 1
			}else {
				right = mid
			}
		}
	}
	return -1
}