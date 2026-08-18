func uniquePaths(m int, n int) int {
    // 2D dynamic programming : bottom up

    // down or right : 2 neighbors
    // start : 0, 0
    // goal : m-1, n-1
    // constraint : fits 32 bit. no need for data type overflow

    // setup
    prevRow := make([]int, n)
    
    // start from bottom row
    for r := m-1; r > -1; r-- {
        // fill the row : use Go zero value
        currRow := make([]int, n)
        // automatically set right column to 1
        // no processing needed
        // specific to this problem
        currRow[n-1] = 1
        // process all columns leftward
        for c := n-2; c > -1; c-- {
            currRow[c] = prevRow[c] + currRow[c+1]
        }
        // move prevRow up
        prevRow = currRow
    }
    // very first cell at 0, 0 = top of the graph = bottom up build completed
    return prevRow[0]
}