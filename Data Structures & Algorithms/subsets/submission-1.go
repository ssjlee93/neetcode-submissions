func subsets(nums []int) [][]int {
	// use backtracking\
	ans := make([][]int, 0)
	path := make([]int, 0)
	// dfs on each node
	var dfs func(int)
	dfs = func(idx int) {
		// Base case: reached end of array
        if idx >= len(nums) {
			temp := make([]int, len(path))
			// slices are references to underlying array
			// without copy, the path slice will be modified
			// and does not guarantee the right set of values
			copy(temp, path)
			ans = append(ans, temp)
			return
		}

		// include curr val to path
		path = append(path, nums[idx])
		dfs(idx+1)
		// exclude curr val to path
		path = path[:len(path)-1]
		dfs(idx+1)
	}

	dfs(0)
	return ans
	
}
