func longestCommonSubsequence(text1 string, text2 string) int {
    // copied solution
    // video : dp bottom up
    m, n := len(text1), len(text2)
    dp := make([][]int, m+1, m+1)
    for row := range m+1 {
        dp[row] = make([]int, n+1, n+1)
    }

    // iterate from goal
    for i := m-1; i > -1; i-- {
        for j := n-1; j > -1; j-- {
            // matched
            if text1[i] == text2[j] {
                // take diagonal + 1 and store
               dp[i][j] = 1 + dp[i+1][j+1]
            } else { // not matched
                // take greater of matched length and store
                dp[i][j] = max(dp[i+1][j], dp[i][j+1])
            }
        }
    }
    return dp[0][0]
}
// time complexity : O n * m
// memory complexity : O n * m
// because this traverses entire 2D grid once.
