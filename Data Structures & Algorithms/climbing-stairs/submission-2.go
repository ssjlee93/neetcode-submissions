func climbStairs(n int) int {
    // memoization
    memo := make(map[int]int)
    // dfs definition.
    // declared in here to inherit n
    var dfs func(i int) int
    dfs = func(i int) int {
        if i >= n {
            if i == n {
                return 1
            }
            return 0
        }
        // skip extra call stacks when you've already been to a layer
        // for example,
        // last 2 branches will always be 1 and 0.
        // with memoization, we can skip 2 call stacks.
        // this happens for every ith lvl.
        // saves a bunch of call stacks
        if val, exists := memo[i]; exists {
            return val
        }
        memo[i] = dfs(i+1) + dfs(i+2)
        return memo[i]
    }
    // use goroutine for dfs
    ch := make(chan int, 1)
    go func() { ch <- dfs(0) }()
    return <-ch
}
