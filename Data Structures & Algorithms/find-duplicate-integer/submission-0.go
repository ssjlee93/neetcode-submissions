func findDuplicate(nums []int) int {
    // fast slow pointer lesson
	// copied solution
	// consider each num as a pointer to an index.
	// fast pointer jumps to the pointer
	// slow pointer iterates
	// when they meet, it means we have a cycle

	// constraint says we will have 1 ~ n
	// 1 guaranteed repeat
	// no 0 in nums
	// meaning fast pointer at 0 doesn't exist
	
	fast := 0
	slow := 0
	for slow != -1 {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}

	slow2 := 0
	for slow != -1 {
		slow = nums[slow]
		slow2 = nums[slow2]
		if slow2 == slow {
			return slow
		}
	}
	return 0
}
