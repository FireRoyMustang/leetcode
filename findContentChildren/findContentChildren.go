package main

import "sort"

// 假设你是一位很棒的家长，想要给你的孩子们一些小饼干。但是，每个孩子最多只能给一块饼干。
// 对每个孩子 i，都有一个胃口值 g[i]，这是能让孩子们满足胃口的饼干的最小尺寸；并且每块饼干 j，
// 都有一个尺寸 s[j] 。如果 s[j] >= g[i]，我们可以将这个饼干 j 分配给孩子 i 。
func main() {

}

func findContentChildren(g []int, s []int) int {
	gLen := len(g)
	sLen := len(s)
	if sLen == 0 || gLen == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	gPointer := 0
	sPointer := 0
	ans := 0
	for {
		if gPointer >= gLen || sPointer >= sLen {
			break
		}
		if g[gPointer] <= s[sPointer] {
			gPointer++
			sPointer++
			ans++
		} else {
			sPointer++
		}
	}
	return ans
}
