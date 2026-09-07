func minSubArrayLen(target int, nums []int) int {
    // sliding window
    ans := len(nums) + 1
    sum := 0 // minimum constraint 1. all sums will be greater than 0
    
    L := 0
    for R := 0; R < len(nums); R++ {
        sum += nums[R]
        
        // check for window constraint
        for sum >= target {
            ans = min(ans, R-L+1)
            sum -= nums[L]
            L++
        }
    }
    if ans > len(nums) {
        return 0
    }
    return ans
}