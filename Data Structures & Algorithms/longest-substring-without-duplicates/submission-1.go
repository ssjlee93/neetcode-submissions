func lengthOfLongestSubstring(s string) int {
	// sliding window
	// reading solution
	ans := 0
	seen := make(map[byte]int)
	l := 0
	for r := 0; r < len(s); r++ {
		if idx, ok := seen[s[r]]; ok {
			l = max(idx+1, l)
		}
		seen[s[r]] = r
		ans = max(ans, r-l+1)
	}
	return ans
}