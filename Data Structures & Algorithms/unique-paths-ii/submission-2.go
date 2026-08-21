func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // topic : dynamic programming - bottom up
	// start at 0, 0
    // goal at m-1, n-1
    // obstacle 1
    // possible unique paths

    // constants
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    dp := make([][]int, m+1, m+1)
    for row := range dp {
        dp[row] = make([]int, n+1, n+1)
    }

    dp[m-1][n-1] = 1

    for i := m-1; i > -1; i-- {
        for j := n-1; j > -1; j-- {
            if obstacleGrid[i][j] == 1 {
                dp[i][j] = 0
            } else {
                dp[i][j] += dp[i+1][j] + dp[i][j+1]
            }
        }
    }
    return dp[0][0]
}