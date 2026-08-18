func uniquePaths(m int, n int) int {
    // 2D dynamic programming

    // down or right : 2 neighbors
    // start : 0, 0
    // goal : m-1, n-1
    // constraint : fits 32 bit. no need for data type overflow

    // memoization (top down) : do recursion then add cache
    cache := make([][]int, m)
    for i := range cache {
        cache[i] = make([]int, n)
    }
    var dfs func(r, c int) int
    dfs = func(r, c int) int {
        // boundary
        if r >= m || c >= n {
            return 0
        }
        // check cache
        if cache[r][c] > 0 {
            return cache[r][c]
        }
        // goal
        if r == m-1 && c == n-1 {
            return 1
        }
        // recursion : go only right or down
        cache[r][c] = dfs(r+1, c) + dfs(r, c+1)
        return cache[r][c]
    }
    return dfs(0, 0)
}