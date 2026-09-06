func numOfSubarrays(arr []int, k int, threshold int) int {
	// sliding window with fixed size
	ans := 0
	windowSum := 0 // use this to calc avg. Saving avg itself is prone to error
	L := 0
	// build initial window
	for i := 0; i < k; i++ {
		windowSum += arr[i]
	}

	avg := windowSum / k
	if avg >= threshold {
		ans++
	}

	for i := k; i < len(arr); i++ {
		n := arr[i]
		windowSum -= arr[L]
		L++
		windowSum += n
		avg = windowSum / k
		if avg >= threshold {
			ans++
		}
	}
	return ans
}
