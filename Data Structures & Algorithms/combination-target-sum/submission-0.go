func combinationSum(nums []int, target int) [][]int {
    // backtracking
    ans := make([][]int, 0)
    path := make([]int, 0)
    var dfs func(int, int)
    dfs = func(idx int, sum int) {
        if sum == target {
            temp := make([]int, len(path))
            copy(temp, path)
            ans = append(ans, temp)
            return
        }
        if idx >= len(nums) || sum > target {
            return
        }

        path = append(path, nums[idx])
        dfs(idx, sum+nums[idx])

        path = path[:len(path)-1]
        dfs(idx+1, sum)
    }

    dfs(0, 0)
    return ans
}
