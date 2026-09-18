func trap(height []int) int {
	// two pointer
	// copied solution
	if len(height) < 3 {
		return 0
	}

	L, R := 0, len(height)-1
	leftMax, rightMax := height[L], height[R]
	ans := 0

	for L < R {
		if leftMax < rightMax {
			L++
			leftMax = max(leftMax, height[L])
			ans += leftMax - height[L]
		} else {
			R--
			rightMax = max(rightMax, height[R])
			ans += rightMax - height[R]
		}
	}
	return ans
}
