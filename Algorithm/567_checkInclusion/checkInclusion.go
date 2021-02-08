package main

// 给定两个字符串 s1 和 s2，写一个函数来判断 s2 是否包含 s1 的排列。
// 换句话说，第一个字符串的排列之一是第二个字符串的子串。
func main() {

}
func checkInclusion(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for _, ch := range s1 {
		cnt[ch-'a']--
	}
	left := 0
	for right, ch := range s2 {
		x := ch - 'a'
		cnt[x]++
		for cnt[x] > 0 {
			cnt[s2[left]-'a']--
			left++
		}
		if right-left+1 == n {
			return true
		}
	}
	return false
}

// func checkInclusion(s1, s2 string) bool {
//     n, m := len(s1), len(s2)
//     if n > m {
//         return false
//     }
//     cnt := [26]int{}
//     for i, ch := range s1 {
//         cnt[ch-'a']--
//         cnt[s2[i]-'a']++
//     }
//     diff := 0
//     for _, c := range cnt[:] {
//         if c != 0 {
//             diff++
//         }
//     }
//     if diff == 0 {
//         return true
//     }
//     for i := n; i < m; i++ {
//         x, y := s2[i]-'a', s2[i-n]-'a'
//         if x == y {
//             continue
//         }
//         if cnt[x] == 0 {
//             diff++
//         }
//         cnt[x]++
//         if cnt[x] == 0 {
//             diff--
//         }
//         if cnt[y] == 0 {
//             diff++
//         }
//         cnt[y]--
//         if cnt[y] == 0 {
//             diff--
//         }
//         if diff == 0 {
//             return true
//         }
//     }
//     return false
// }
