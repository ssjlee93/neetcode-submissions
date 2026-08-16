func rob(nums []int) int {
    // copied solution
    // dynamic programming

    // memoization
    cache := make(map[int]int)

    // DFS
    var dfs func(i int) int
    dfs = func(i int) int {
        // base case
        if i >= len(nums) {
            return 0
        }
        if val, ok := cache[i]; ok {
            return val
        }
        cache[i] = max(dfs(i+1), nums[i]+dfs(i+2))
        return cache[i]
    }

    return dfs(0)
}
