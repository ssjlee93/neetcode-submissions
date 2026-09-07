func lengthOfLongestSubstring(s string) int {
	// sliding window
	// watching solution
	ans := 0
	seen := make(map[byte]bool)
	l := 0
	for r := 0; r < len(s); r++ {
		for seen[s[r]] {
			delete(seen, s[l])
			l++
		}
		seen[s[r]] = true
		ans = max(ans, r-l+1)
	}
	return ans
}