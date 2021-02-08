package main

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
