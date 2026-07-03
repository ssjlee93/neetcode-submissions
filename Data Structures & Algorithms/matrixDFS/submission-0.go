type Point struct {
    X int
    Y int
}

func countPaths(grid [][]int) int {
    ROWS, COLS := len(grid), len(grid[0])
    visited := make(map[Point]bool)
    var dfs func(int, int, map[Point]bool) int
    dfs = func (r, c int, visit map[Point]bool) int {
        point := Point{c, r}
        // base case : boundary
        if r < 0 || c < 0 || r == ROWS || c == COLS || 
        visit[point] || 
        grid[r][c] == 1 {
            return 0
        }

        // goal
        if r == ROWS-1 && c == COLS-1 {
            return 1
        }

        // mark visited
        visit[point] = true
        
        count := 0
        count += dfs(r+1, c, visit)
        count += dfs(r-1, c, visit)
        count += dfs(r, c+1, visit)
        count += dfs(r, c-1, visit)

        // backtrack
        // allow other DFS to find a path by re-visited a point
        visit[point] = false
        return count
    }
    return dfs(0, 0, visited)
}
