func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums)*2, len(nums)*2)
    
    j := len(nums)
    for i := 0; i < len(nums); i++ {
        ans[i] = nums[i]
        ans[j] = nums[i]
        j++
    }
    return ans
}
