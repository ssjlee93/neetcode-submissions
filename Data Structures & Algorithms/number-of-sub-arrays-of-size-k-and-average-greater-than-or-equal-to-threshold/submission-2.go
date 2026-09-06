func numOfSubarrays(arr []int, k int, threshold int) int {
	// sliding window with fixed size
	// AI helped. old one i learned uses i-k
	ans := 0
	windowSum := 0
	// build the window
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	// check initial window validity
	avg := windowSum / k
	if avg >= threshold {
		ans++
	}
	// loop from end of window
	for i := k; i < len(arr); i++ {
		// add one and remove one each iter
		windowSum += arr[i]
		windowSum -= arr[i-k]
		avg := windowSum / k
		if avg >= threshold {
			ans++
		}
	}
	return ans
	
}