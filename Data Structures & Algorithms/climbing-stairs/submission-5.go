func climbStairs(n int) int {
    // top down dynamic programming
    // coding based on "algorithm" section

    // Create a cache array cache of size n initialized with -1.
    cache := make(map[int]int)
    // Define a recursive function dfs(i):
    var dfs func(i int) int
    dfs = func(i int) int {
        // If i == n, return 1 (valid way).
        if i == n {
            return 1
        }
        // If i > n, return 0 (invalid path).
        if i > n {
            return 0
        }
        // If cache[i] != -1, return the cached value.
        if val, ok := cache[i]; ok {
            return val
        }
        // Otherwise:
        // Compute dfs(i + 1) + dfs(i + 2)
        computed := dfs(i+1) + dfs(i+2)
        // Store the result in cache[i]
        cache[i] = computed
        return computed
    }
    // Return dfs(0) as the final answer.
    return dfs(0)
}
