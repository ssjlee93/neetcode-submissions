func removeDuplicates(nums []int) int {
    // modify nums in-place
    // two pinter
    // a pointer for insert
    // another pointer to iterate

    // algorithm : 
    // since array is sorted, 
    // we can a duplicate by checking current max
    // if curr max == num, then it's a duplicate
    // when we find a duplicate, use a flag to check if it's second visit.
    // second visits and beyond will be ignored and move right pointer
    // first visit will perform a swap
    // if curr max < num, then it's a new num
    // duplicates will move the insert pointer
    // new num will perform swap
    if len(nums) <= 2 {
        return len(nums)
    }
    insert := 0
    currMax := -10001   // from constraint
    doubled := false
    for i := 0; i < len(nums); i++ {
        n := nums[i]
        if currMax < n {
            currMax = n
            nums[insert] = n
            insert++
            doubled = false
        } else {
            // since array is sorted,
            // only case : currMax == n
            if !doubled {
                doubled = true
                nums[insert] = n
                insert++
            }
        }
    }
    return insert
}