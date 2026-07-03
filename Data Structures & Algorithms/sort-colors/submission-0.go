func sortColors(nums []int) {
    // bucket sort
    // Step 1: Count occurrences
    count := [3]int{}
    for _, num := range nums {
        count[num]++
    }
    
    // Step 2: Calculate boundaries
    // 0s: [0, count[0])
    // 1s: [count[0], count[0] + count[1])
    // 2s: [count[0] + count[1], n)
    
    // Step 3: Place elements in their correct positions
    i := 0
    for color := 0; color < 3; color++ {
        for j := 0; j < count[color]; j++ {
            nums[i] = color
            i++
        }
    }
    
    return 
}
