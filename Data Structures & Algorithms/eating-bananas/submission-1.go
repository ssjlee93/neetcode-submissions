func minEatingSpeed(piles []int, h int) int {
	// find max value
    max := 0
    for _, pile := range piles {
        if pile > max {
            max = pile
        }
    }

    // binary search for the rate of consumption k
    low, high := 1, max

    for low <= high {
        // find mid rate k
        mid := low + (high-low)/2

        totalTime := calculateTime(piles, mid)
        
        // compare against the allotted time
        if totalTime <= h {
            high = mid-1
        } else {
            low = mid+1
        }
    }
    return low
}

func calculateTime(piles []int, k int) int {
    totalTime := 0
    for _, pile := range piles {
        hours := (pile + k-1) / k
        totalTime += hours
    }
    return totalTime
}
