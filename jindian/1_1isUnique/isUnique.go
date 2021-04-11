package main

// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

func main() {

}

// 不使用额外空间：位运算
func isUnique(astr string) bool {
	if len(astr) <= 1 {
		return true
	}
	flag := 0
	for i := 0; i < len(astr); i++ {
		bit := int(astr[i] - 'a')
		num := 1 << bit
		if flag&num != 0 {
			return false
		} else {
			flag = flag | num
		}

	}
	return true
}

// func isUnique(astr string) bool {
// 	if len(astr) <= 1 {
// 		return true
// 	}
// 	set := make(map[byte]bool)
// 	for i := 0; i < len(astr); i++ {
// 		if set[astr[i]] {
// 			return false
// 		}
// 		set[astr[i]] = true
// 	}
// 	return true
// }
