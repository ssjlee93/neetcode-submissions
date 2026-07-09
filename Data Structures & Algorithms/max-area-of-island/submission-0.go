type Pair struct {
    Row int
    Col int
}

func maxAreaOfIsland(grid [][]int) int {
    // use DFS to scan for each island.
    // more intuitive to scan each island,
    // log its maximum,
    // and search for another.

    // constants
    ROWS, COLS := len(grid), len(grid[0])
    ans := 0
    visited := make(map[Pair]bool) // keep track of visited

    var dfs func(int, int, map[Pair]bool) int
    dfs = func(r, c int, visited map[Pair]bool) int {
        pair := Pair{r, c}
        // boundary
        if r < 0 || c < 0 || r == ROWS || c == COLS ||
        visited[pair] ||
        grid[r][c] == 0 {
            return 0
        }

        // only handle 1s
        visited[pair] = true

        count := 1
        count += dfs(r+1, c, visited)
        count += dfs(r-1, c, visited)
        count += dfs(r, c+1, visited)
        count += dfs(r, c-1, visited)
        
        return count
    }
    
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == 1 && !visited[Pair{r, c}] {
                area := dfs(r, c, visited)
                if area > ans {
                    ans = area
                }
            }
        }
    }
    
    return ans
}