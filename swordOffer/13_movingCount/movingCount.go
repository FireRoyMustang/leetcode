package main

// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始
// 移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之
// 和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入
// 方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？
var limit, x, y int
var visited [][]bool

func movingCount(m int, n int, k int) int {
	limit = k
	x = m
	y = n
	visited = make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	return dfs(0, 0)
}
func dfs(m, n, lastSum int) int {
	if m >= x || n >= y || lastSum > limit || visited[m][n] {
		return 0
	}
	visited[m][n] = true
	xFlag, yFlag := false, false
	if m%10 == 9 {
		xFlag = true
	}
	if n%10 == 9 {
		yFlag = true
	}
	ans := 1
	if xFlag {
		ans += dfs(m+1, n, lastSum-8)
	} else {
		ans += dfs(m, n, lastSum+1)
	}
	if yFlag {
		ans += dfs(m, n+1, lastSum-8)
	} else {
		ans += dfs(m, n, lastSum+1)
	}
	return ans
}
