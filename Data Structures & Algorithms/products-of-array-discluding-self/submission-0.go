func productExceptSelf(nums []int) []int {
	// prefix sum lesson
	// help from Jaehak

	// make prefix product
	// make suffix product
	// to get ans for each position, 
	// prefix[i-1] * suffix[i+1]

	prefix := make([]int, len(nums)+1, len(nums)+1)
	suffix := make([]int, len(nums)+1, len(nums)+1)

	prefix[0] = 1
	suffix[len(nums)] = 1

	i := 0
	j := len(nums)-1
	for i < len(nums) {
		prefix[i+1] = prefix[i] * nums[i]
		suffix[j] = suffix[j+1] * nums[j]
		i++
		j--
	}

	ans := make([]int, 0, len(nums))
	for i := 1; i < len(prefix); i++ {
		left := prefix[i-1]
		right := suffix[i]
		product := left * right
		ans = append(ans, product)
	}
	return ans
}
