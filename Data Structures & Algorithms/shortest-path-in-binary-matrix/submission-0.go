type Pair struct {
	Row int
	Col int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	// constants for the problem
	ROWS, COLS := len(grid), len(grid[0])
	if grid[0][0] != 0 || grid[ROWS-1][COLS-1] != 0 {
		return -1
	}
	ans := 1
	// vars for BFS
	visited := make(map[Pair]bool)
	q := make([]Pair, 0)
	// starting point
	origin := Pair{0, 0}
	q = append(q, origin)
	visited[origin] = true
	// define helper function
	var processNeighbor func(int, int)
	processNeighbor = func(r, c int) {
		// check boundary
		point := Pair{r, c}
		if r < 0 || c < 0 || r >= ROWS || c >= COLS ||
			visited[point] ||
			grid[point.Row][point.Col] != 0 { // condition specific to the problem
			return
		}
		q = append(q, point)
		visited[point] = true
		return
	}

	// BFS
	for len(q) > 0 {
		lenLvl := len(q)
		for range lenLvl {
			// pop
			p := q[0]
			q = q[1:]
			r, c := p.Row, p.Col

			// goal
			if r == ROWS-1 && c == COLS-1 {
				return ans
			}

			// breadth search
			processNeighbor(r+1, c)
			processNeighbor(r-1, c)
			processNeighbor(r, c+1)
			processNeighbor(r, c-1)
			processNeighbor(r+1, c+1)
			processNeighbor(r-1, c-1)
			processNeighbor(r+1, c-1)
			processNeighbor(r-1, c+1)
		}
		ans++
	}

	return -1
}
