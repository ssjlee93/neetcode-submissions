func maxSubarraySumCircular(nums []int) int {
    // Kadane's algorithm
    // copied solution
    globalMax, globalMin := nums[0], nums[0]
    currMax, currMin := 0, 0
    total := 0

    for _, n := range nums {
            currMax = max(currMax+n, n)
            currMin = min(currMin+n, n)

        total += n
        if currMax > globalMax {
            globalMax = currMax
        }
        if currMin < globalMin {
            globalMin = currMin
        }
    }
    // when globalMax is negative, even though we subtract globalMin, we still end up with neg.
    // in fact, we get positive num.
    // so we return just globalMax
    if globalMax < 0 {
        return globalMax
    }
    if total - globalMin > globalMax {
        return total - globalMin
    } 
    return globalMax
}
