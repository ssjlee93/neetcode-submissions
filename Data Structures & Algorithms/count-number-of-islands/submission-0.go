type Point struct {
    Row int
    Col int
}

func numIslands(grid [][]byte) int {
    ROWS, COLS := len(grid), len(grid[0])
    // use DFS'
    visited := make(map[Point]bool)
    var dfs func(int, int, map[Point]bool)
    dfs = func(r, c int, visited map[Point]bool) {
        point := Point{r, c}
        // boundary
        if r < 0 || c < 0 || r == ROWS || c == COLS ||
        visited[point] ||
        grid[r][c] == '0' {
            return
        }
    // we only recurse when we have a 1
    // add each visited node to visited
        visited[point] = true

        dfs(r+1, c, visited)
        dfs(r-1, c, visited)
        dfs(r, c+1, visited)
        dfs(r, c-1, visited)
    }

    islands := 0
    for r := 0; r < ROWS; r++ {
        for c := 0; c < COLS; c++ {
            if grid[r][c] == '1' && !visited[Point{r, c}] {
                dfs(r, c, visited)
                islands++
            }
        }
    }
    return islands
}
