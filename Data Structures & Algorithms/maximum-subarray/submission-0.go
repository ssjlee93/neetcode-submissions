func maxSubArray(nums []int) int {
    // Kadane's algorithm
    ans := -10001 // taken from constraint
    curr := 0

    for _, n := range nums {
        // negative num resets the window
        if curr < 0 {
            curr = 0
        }
        curr += n
        if curr > ans {
            ans = curr
        }
    }
    return ans
}

// time complexity : O(n)
// space complexity : O(1)
