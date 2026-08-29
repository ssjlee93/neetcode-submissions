func countBits(n int) []int {
    // copied solution
    // Brian Kernighan’s Algorithm
    ans := make([]int, n+1, n+1)
    for i := 1; i <= n; i++ {
        num := i
        for num != 0 {
            ans[i]++
            // main action
            // removes the last 1 in the number
            // subtracting 1 means breaking the first one from the right (right most 1)
            // and bitwise & will match each position to be AND
            // 0010 1000 & 0010 0111 -> 0010 0000
            // the 1000 breaks down to 0111 and each position can't match, so it becomes 0000
            // if we have 10010, then 10010 & 10001 -> 10000
            // always clears right most 1
            num &= (num - 1)
        }

    }
    return ans
}
