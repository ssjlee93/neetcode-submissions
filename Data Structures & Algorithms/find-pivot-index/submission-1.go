func pivotIndex(nums []int) int {
	// prefix sum lesson
    
    // make prefix sum
    // initialize with 0 as initial value
    // loop through and check if last sum - prefix[i] = prefix[i-1]

    // o(n) time
    // o(n) space

    prefix := make([]int, 0, len(nums)+1)
    prefix = append(prefix, 0)
    
    // form prefix sum
    for _, n := range nums {
        prefix = append(prefix, prefix[len(prefix)-1] + n)
    }
    
    // loop and process
    ans := -1
    for i := 1; i < len(prefix); i++ {
        if prefix[len(prefix)-1] - prefix[i] == prefix[i-1] {
            return i-1
        }
    }
    return ans
}
