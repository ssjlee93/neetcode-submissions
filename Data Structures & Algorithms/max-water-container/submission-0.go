func maxArea(heights []int) int {
	// two pointer
	// track Left and Right
	// calculate width
	// move pointer for the smaller one
	// track max size 

	L, R := 0, len(heights)-1
	ans := 0
	
	for L < R {
		left := heights[L]
		right := heights[R]
		
		// take smaller number as height
		height := left
		if right < left {
			height = right
		}
		
		// calculate
		width := R - L
		volume := width * height
		if volume > ans {
			ans = volume
		}

		// move smaller pointer
		if height == left {
			L++
		} else {
			R--
		}
	}

	return ans
}
