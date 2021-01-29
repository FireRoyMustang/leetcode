package main

import "fmt"

func main() {
	fmt.Println(isMatch("aa", "a*"))
}
func isMatch(s string, p string) bool {
	m, n := len(s)+1, len(p)+1
	dp := make([][]bool, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]bool, n)
	}
	dp[0][0] = true
	for i := 2; i < n; i += 2 {
		dp[0][i] = dp[0][i-2] && p[i-1] == '*'
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if p[j-1] == '*' {
				// dp[i][j-1]：出现1次  dp[i][j-2]：出现0次 dp[i-1][j]：出现大于1次
				dp[i][j] = dp[i][j-1] || dp[i][j-2] || dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.')
			} else {
				dp[i][j] = dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[m-1][n-1]
}

func isMatch2(s string, p string) bool {
	m, n := len(s), len(p)
	matches := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	f := make([][]bool, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]bool, n+1)
	}
	f[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				f[i][j] = f[i][j] || f[i][j-2]
				if matches(i, j-1) {
					f[i][j] = f[i][j] || f[i-1][j]
				}
			} else if matches(i, j) {
				f[i][j] = f[i][j] || f[i-1][j-1]
			}
		}
	}
	return f[m][n]
}
