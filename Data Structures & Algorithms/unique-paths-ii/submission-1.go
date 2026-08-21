func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // topic : dynamic programming - bottom up
	// start at 0, 0
    // goal at m-1, n-1
    // obstacle 1
    // possible unique paths

    // constants
    m, n := len(obstacleGrid), len(obstacleGrid[0])
    prevRow := make([]int, n, n)

    for i := m-1; i > -1; i-- {
        currRow := make([]int, n, n)
        for j := n-1; j > -1; j-- {
            if obstacleGrid[i][j] == 0 {
                if i == m-1 && j == n-1 {
                    currRow[j] = 1
                    continue
                }
                right := 0
                if j + 1 < n {
                    right = currRow[j+1]
                }
                currRow[j] = right + prevRow[j]
            }
        }
        prevRow = currRow
    }
    return prevRow[0]
}