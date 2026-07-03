func removeElement(nums []int, val int) int {
    // return if nums doesn't have enough
    if len(nums) == 0 {
        return 0
    }
    left := 0
    // mvoe through and when it matches val, shift one to left
    for right := 0; right < len(nums); right++ {
        if nums[right] == val {
            continue
        }
        nums[left] = nums[right]
        left++
    }
    return left
}
