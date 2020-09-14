package level4

// 判断子序列
func isSubsequence(s string, t string) bool {
	son, father := len(s), len(t)
	i , j := 0, 0
	for i < son && j < father {
		if s[i] == t[j]{
			i ++
		}
		j ++
	}
	return son == i
}

func uniquePaths(m int, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}

	dp := make([][]int, 0)

	for i := 0; i < m; i ++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 0; i < n; i ++ {
		dp[0][i] = 1
	}

	for i := 1; i < m ; i ++ {
		for j := 1; j < n ; j ++ {
			dp[i][j] = dp[i-1][j] + dp[i][j -1]
		}
	}
	return dp[m-1][n-1]
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid == nil || len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0{
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	var dp [][]int
	for i := 0; i < m ; i ++ {
		tmp := make([]int, n)
		for j := 0 ; j < n ; j ++ {
			tmp = append(tmp, 0)
		}
		dp = append(dp, tmp)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0 ; i ++ {
		dp[i][0] = 1
	}

	for j := 0; j < n && obstacleGrid[0][j] == 0 ; j ++ {
		dp[0][j] = 1
	}

	for i := 1; i < m ; i ++ {
		for j := 1; j < n ; j ++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j -1]
			}
		}
	}
	return dp[m-1][n-1]
}