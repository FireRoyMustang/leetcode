package main

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。
// 请你找出并返回这两个正序数组的 中位数 。
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	totLen := len1 + len2
	mid := totLen / 2
	if totLen%2 == 1 {
		return float64(getKthNum(nums1, nums2, mid+1))
	}
	return float64(getKthNum(nums1, nums2, mid)+getKthNum(nums1, nums2, mid+1)) / 2
}

// k代表第i个数而不是坐标为i
func getKthNum(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	len1 := len(nums1)
	len2 := len(nums2)
	for {
		if index1 == len1 {
			return nums2[index2+k-1]
		}
		if index2 == len2 {
			return nums1[index1+k-1]
		}
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1+half, len1) - 1
		newIndex2 := min(index2+half, len2) - 1
		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 <= pivot2 {
			k -= newIndex1 - index1 + 1
			index1 = newIndex1 + 1
		} else {
			k -= newIndex2 - index2 + 1
			index2 = newIndex2 + 1
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
