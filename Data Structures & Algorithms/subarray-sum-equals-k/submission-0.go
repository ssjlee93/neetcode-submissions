func subarraySum(nums []int, k int) int {
	// prefix sum lesson
	// idea C jaehak
	// and copied solution
	
	// prefix hashmap sum 
	// hashmap key 0 == ans
	// hashmap k - prefix sum and if hashmap key ok, then ans + the key

	prefix := make(map[int]int)
	prefix[0] = 1
	currSum := 0
	ans := 0

	for _, n := range nums {
		currSum += n
		diff := currSum - k

		ans += prefix[diff]
		prefix[currSum]++
	}
	return ans
}
