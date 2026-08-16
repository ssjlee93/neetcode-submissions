func rob(nums []int) int {
    // copied video solution
    // dynamic programming bottom up with only 2 last necessary values
    rob1, rob2 := 0, 0

    for _, n := range nums {
        temp := max(rob2, rob1 + n)
        rob1 = rob2
        rob2 = temp
    }
    return rob2
}
