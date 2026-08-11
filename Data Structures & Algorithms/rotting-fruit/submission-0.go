type Pair struct {
    Row int
    Col int
}

func orangesRotting(grid [][]int) int {
    // constants
    ROWS, COLS := len(grid), len(grid[0])
    ans := 0
    // BFS vars
    q := make([]Pair, 0)
    visited := make(map[Pair]bool)

    var processNeighbor func(int, int)
    processNeighbor = func(r, c int) {
        point := Pair{r, c}
        // boundaries
        if r < 0 || c < 0 || r >= ROWS || c >= COLS || visited[point] {
            return
        }
        visited[point] = true
        if grid[r][c] != 1 {
            return
        }
        grid[r][c] = 2
        q = append(q, point)
        return
    }
    // 1. scan the grid for a rotten fruit
    // this will be our first lvl of starting point
    for i := range ROWS {
        for j := range COLS {
            if grid[i][j] == 2 {
                q = append(q, Pair{i, j})
                visited[Pair{i, j}] = true
            }
        }
    }
    // 2. do a BFS from the rotten fruit
    for len(q) > 0 {
        lvl := len(q)
        // tracks if an orange has changed.
        // a lvl with all rotten or empty doesn't convert anything
        foundFresh := false 
        for range lvl {
            // pop
            p := q[0]
            q = q[1:]
            r, c := p.Row, p.Col

            // expand breadth
            oldLen := len(q)
            processNeighbor(r+1, c)
            processNeighbor(r-1, c)
            processNeighbor(r, c+1)
            processNeighbor(r, c-1)
            if len(q) > oldLen {
                foundFresh = true
            }
        }
        // each lvl means 1 extra min
        if foundFresh {
            ans++
        }
    }
    
    // 3. scan the grid if any fresh orange remains
    for i := range ROWS {
        for j := range COLS {
            if grid[i][j] == 1 {
                return -1
            }
        }
    }
    return ans
}