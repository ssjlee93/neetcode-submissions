func hasDuplicate(nums []int) bool {
    // would use hashset for counting
    seen := make(map[int]bool)
    
    for _, n := range nums {
        _, ok := seen[n]
        if ok {
            return true
        }
        seen[n] = true
    }

    return false
}
