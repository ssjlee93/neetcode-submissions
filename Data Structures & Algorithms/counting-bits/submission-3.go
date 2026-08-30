func countBits(n int) []int {
    // copied solution
    // dynamic programming bottom up
    
    // 1. analyze the pattern of the bits
    // 0000 -> 0
    // 0001 -> 1
    // 0010 -> 1 = 1 + i-2
    // 0011 -> 2 = 1 + i-2
    // 0100 -> 1 = 1 + i-4
    // 0101 -> 2 = 1 + i-4
    // 0110 -> 2 ...
    // 0111
    // 1000
    // 2. based on the pattern, use it to solve the sequence of problems
    // 0 and 1 pattern repeats every time a new 1 is introduced to the left.
    // 3. check for any conditions to change pattern
    // every time i hits power of 2, the pattern triggers

    // initialize with zero value to account for first 0
    // also can add 0 as first value
    dp := make([]int, n+1, n+1)
    // offset drives the i that hits power of 2 and trigger pattern
    // cannot be 0 since that will not calculate power of 2
    offset := 1

    for i := 1; i < n+1; i++ {
        // check if i hit offset 
        if offset * 2 == i {
            offset = i
        }
        // form dp
        dp[i] = 1 + dp[i-offset]
    }
    return dp
}
