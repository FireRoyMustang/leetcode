package main

// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。

func main() {

}

func sumNums(n int) int {
	ans := 0
	var sum func(int) bool
	sum = func(n int) bool {
		ans += n
		return n > 0 && sum(n-1)
	}
	sum(n)
	return ans
}

// func sumNums(n int) int {
//     ans, A, B := 0, n, n + 1
//     addGreatZero := func() bool {
//         ans += A
//         return ans > 0
//     }

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1

//     _ = ((B & 1) > 0) && addGreatZero()
//     A <<= 1
//     B >>= 1
//     return ans >> 1
// }

// 作者：LeetCode-Solution
// 链接：https://leetcode-cn.com/problems/qiu-12n-lcof/solution/qiu-12n-by-leetcode-solution/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
