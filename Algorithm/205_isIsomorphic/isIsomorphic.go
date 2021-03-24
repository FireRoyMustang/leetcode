package main

// 给定两个字符串 s 和 t，判断它们是否是同构的。
// 如果 s 中的字符可以按某种映射关系替换得到 t ，那么这两个字符串是同构的。
// 每个出现的字符都应当映射到另一个字符，同时不改变字符的顺序。
// 不同字符不能映射到同一个字符上，相同字符只能映射到同一个字符上，字符可以映射到自己本身。

// 可以假设 s 和 t 长度相同。
import "fmt"

func main() {
	fmt.Println(isIsomorphic("ab", "aa"))
	fmt.Println(isIsomorphic("ab", "ca"))
	fmt.Println(isIsomorphic("13", "42"))
}
func isIsomorphic(s string, t string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	s2t := make([]byte, 128)
	t2s := make([]byte, 128)
	for index := range s {
		chars, chart := s[index], t[index]
		// 因为一一映射，所以一定要这样
		if s2t[chars] > 0 && s2t[chars] != chart || t2s[chart] > 0 && t2s[chart] != chars {
			return false
		}
		s2t[chars] = chart
		t2s[chart] = chars
	}
	return true
}

func isIsomorphic2(s, t string) bool {
	s2t := map[byte]byte{}
	t2s := map[byte]byte{}
	for i := range s {
		x, y := s[i], t[i]
		if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
			return false
		}
		s2t[x] = y
		t2s[y] = x
	}
	return true
}
