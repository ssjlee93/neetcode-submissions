func longestCommonSubsequence(text1 string, text2 string) int {
    // copied solution
    // dp top down
    m, n := len(text1), len(text2)
    cache := make(map[[2]int]int)
    
    var dfs func(i, j int) int
    dfs = func(i, j int) int {
        point := [2]int{i, j}
        // base case
        if i >= m || j >= n {
            return 0
        }
        // memoization
        if val, ok := cache[point]; ok {
            return val
        }
        // goal
        if text1[i] == text2[j] {
            cache[point] = 1 + dfs(i+1, j+1)
        } else {
            cache[point] = max(dfs(i+1, j), dfs(i, j+1))
        }
        return cache[point]
    }
    return dfs(0, 0)
}
// time complexity : O n * m
// memory complexity : O n * m
// because this traverses entire 2D grid once.
