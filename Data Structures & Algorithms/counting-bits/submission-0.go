func countBits(n int) []int {
	ans := make([]int, 0)
	for i := 0; i < n+1; i++ {
		ans  = append(ans, countOne(i))
	}
	return ans
}

func countOne(n int) int {
	count := 0
	for n > 0 {
		if n & 1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}
