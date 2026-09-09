func maxTurbulenceSize(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	ans := 1
	currPattern := 0
	l := 0
	for r := 0; r < len(arr)-1; r++ {
		pattern := determinePattern(arr[r], arr[r+1], r)
		if pattern == 0 {
			currPattern = 0
			l = r + 1
			continue
		}
		if pattern != currPattern {
			l = r
			currPattern = pattern
		}
		ans = max(ans, r+1-l+1)
	}
	return ans
}

func determinePattern(x, y, i int) int {
	if i % 2 == 0 && x < y {
		return 1
	} else if i % 2 != 0 && x > y {
		return 1
	} else if i % 2 == 0 && x > y {
		return 2
	} else if i % 2 != 0 && x < y {
		return 2
	}
	return 0
}