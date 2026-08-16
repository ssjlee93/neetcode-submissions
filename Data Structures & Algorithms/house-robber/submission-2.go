func rob(nums []int) int {
    // copied solution
    // dynamic programming bottom up

    // edge case
    n := len(nums)
    if n == 0 {
        return 0
    }
    if n == 1 {
        return nums[0]
    }

    dp := make([]int, n)
    // initial values
    dp[0] = nums[0]
    dp[1] = max(nums[0], nums[1])

    // iterate to rob
    for i := 2; i < n; i++ {
        dp[i] = max(dp[i-1], dp[i-2]+nums[i])
    }

    return dp[n-1]
}
