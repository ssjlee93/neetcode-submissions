func subsets(nums []int) [][]int {
	// use backtracking\
	ans := make([][]int, 0)
	path := make([]int, 0)
	// dfs on each node
	var dfs func(int)
	dfs = func(idx int) {
		
		temp := make([]int, len(path))
		// slices are references to underlying array
		// without copy, the path slice will be modified
		// and does not guarantee the right set of values
		copy(temp, path)
		ans = append(ans, temp)

		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i+1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
	
}
