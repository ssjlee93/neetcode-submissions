func climbStairs(n int) int {
    // bottom up dynamic programming
    // coded based on "algorithm" in solution

    // If n <= 2, return n.
    // edge case : cannot process
    if n <= 2 {
        return n
    }
    // Create a DP array where dp[i] = number of ways to reach step i.
    dp := make([]int, n+1, n+1)
    // Initialize:
    dp[1] = 1
    dp[2] = 2
    // For i from 3 to n:
    for i := 3; i <= n; i++ {
        dp[i] = dp[i - 1] + dp[i - 2]
    }
    return dp[n]
}
