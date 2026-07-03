func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    insertIdx := 0
    for i := 1; i < len(nums); i++ {
        curr := nums[i]
        if nums[insertIdx] == curr {
            continue
        }
        
        insertIdx++
        nums[insertIdx] = curr
    }
    insertIdx++
    return insertIdx
}
