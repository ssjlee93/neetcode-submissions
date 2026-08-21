func uniquePathsWithObstacles(obstacleGrid [][]int) int {
    // topic : dynamic programming - memoization
	// start at 0, 0
    // goal at m-1, n-1
    // obstacle 1
    // possible unique paths
    // BFS with backtracking
    // DFS with memoization = dynamic programming

    // constants
    m, n := len(obstacleGrid), len(obstacleGrid[0])

    // DFS
    cache := make(map[[2]int]int)
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        point := [2]int{ r, c }
        // boundary check
        if r < 0 || c < 0 || r >= m || c >= n ||
        obstacleGrid[r][c] == 1 {
            return 0
        }
        // cache check
        if val, ok := cache[point]; ok {
            return val
        }
        // goal
        if r == m-1 && c == n-1 {
            return 1
        }
        // recurse
        cache[point] = dfs(r+1, c) + dfs(r, c+1)
        return cache[point]
    }

    return dfs(0, 0)
}
