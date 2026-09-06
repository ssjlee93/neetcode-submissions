func containsNearbyDuplicate(nums []int, k int) bool {
    // sliding window with fixed size
    // initialzie L
    L := 0
    // initialize hashset to keep track of nums in the window
    window := make(map[int]bool)
    // loop through R
    for R, n := range nums {
        //  if window exceeds,
        if R - L > k {
            delete(window, nums[L])
            // increment L
            L++
        }
        // check if hashset R exists
        if ok := window[n]; ok {
            // immediately return true
            return true
        }

        // add num R into hashset
        window[n] = true
    }
    
    return false
}
