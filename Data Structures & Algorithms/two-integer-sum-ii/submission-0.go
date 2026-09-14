func twoSum(numbers []int, target int) []int {
	// two pointer
	// array already sorted
	L, R := 0, len(numbers)-1
	for L < R {
		sum := numbers[L] + numbers[R]
		if sum == target {
			return []int{L+1, R+1}
		}

		if sum > target {
			R--
		} else if sum < target {
			L++
		}
	}
	return []int{}
}
