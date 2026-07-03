func combinationSum(nums []int, target int) [][]int {
    // backtracking
    ans := make([][]int, 0)
    path := make([]int, 0)
    sum := 0
    var dfs func(int)
    dfs = func(idx int) {
        if idx >= len(nums) || sum > target {
            return
        }
        if sum == target {
            temp := make([]int, len(path))
            copy(temp, path)
            ans = append(ans, temp)
            return
        }

        path = append(path, nums[idx])
        sum += nums[idx]
        dfs(idx)

        path = path[:len(path)-1]
        sum -= nums[idx]
        dfs(idx+1)
    }

    dfs(0)
    return ans
}
