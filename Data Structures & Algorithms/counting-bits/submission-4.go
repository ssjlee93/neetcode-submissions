func countBits(n int) []int {
    // copied solution
    // dynamic programming bottom up with bitwise operation
    
    // 1. analyze the pattern of the bits
    // 0000 -> 0
    // 0001 -> 1
    // 0010 -> 1 = earlier cycle[1] + 1
    // 0011 -> 2 = earlier cycle[1] + 1
    // 0100 -> 1 = earlier cycle[2] + 0
    // 0101 -> 2 = earlier cycle[2] + 1
    // 0110 -> 2 = earlier cycle[3] + 0
    // 0111 -> 3 = earlier cycle[3] + 1
    // 1000 -> 1
    // 2. based on the pattern, use it to solve the sequence of problems
    // every time a new 1 is introduced to the left, it refers back to previous cycle + 1
    // (how the f do you come up with this pattern...)
    // 3. check for any conditions to change pattern
    // every time i hits power of 2, the pattern triggers

    // initialize with zero value to account for first 0
    // also can add 0 as first value
    dp := make([]int, n+1, n+1)

    for i := 1; i < n+1; i++ {
        dp[i] = dp[i >> 1] + (i & 1)
    }
    return dp
}
