package tools

func getNext(s string) []int {
	m := len(s)

	if m == 0 {
		return nil
	}
	next := make([]int, m+1)
	for i, j := 2, 0; i <= m; i++ {
		// 匹配不成功的话，j = next[j]
		for j > 0 && s[i] != s[j+1] {
			j = next[j]
		}
		// 匹配成功的话，先让j++
		if s[i] == s[j+1] {
			j++
		}
		// 更新next[i]
		next[i] = j
	}
	return next

}
