package main

// 给你一个整数 n ，请你找出并返回第 n 个 丑数 。

// 丑数 就是只包含质因数 2、3 和/或 5 的正整数。
func main() {

}
func nthUglyNumber(n int) int {
	uglys := make([]int, n)
	uglys[0] = 1
	ptr2, ptr3, ptr5 := 0, 0, 0
	for i := 0; i < n; i++ {
		ugly := min(uglys[ptr2]*2, uglys[ptr3]*3, uglys[ptr5]*5)
		uglys[i] = ugly
		if ugly%2 == 0 {
			ptr2++
		}
		if ugly%3 == 0 {
			ptr3++
		}
		if ugly%5 == 0 {
			ptr5++
		}
	}
	return uglys[n-1]
}

func min(i, j, k int) int {
	if i < j {
		if k < i {
			return k
		}
		return i
	}
	if k < j {
		return k
	}
	return j
}
