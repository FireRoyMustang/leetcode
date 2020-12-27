package main

import "strings"

func main() {

}
func replaceSpace(s string) string {
	ans := ""

	ans = strings.Replace(s, " ", "%20", -1)

	return ans
}
