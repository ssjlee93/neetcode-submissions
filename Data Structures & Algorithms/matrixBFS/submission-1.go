type Pair struct {
	Row int
	Col int
}

func shortestPath(grid [][]int) int {
	// constants
	ROWS, COLS := len(grid), len(grid[0])
	ans := 0

	// BFS variables
	q := make([]Pair, 0)
	visited := make(map[Pair]bool)

	var processNeighbor func(int, int)
	processNeighbor = func(r, c int) {
		point := Pair{r, c}

		// boundary
		if r < 0 || c < 0 || r == ROWS || c == COLS ||
		visited[point] ||
		grid[r][c] == 1 {
			return
		}

		// process our queue
		q = append(q, point)
		visited[point] = true
	}

	// add starting point
	start := Pair{0, 0}
	q = append(q, start)
	visited[start] = true

	// BFS
	for len(q) > 0 {
		length := len(q)
		for _ = range length {
			// pop
			pair := q[0]
			q = q[1:]
			r, c := pair.Row, pair.Col

			// goal
			if r == ROWS-1 && c == COLS-1 {
				return ans
			}

			// go breadth
			processNeighbor(r+1, c)
			processNeighbor(r-1, c)
			processNeighbor(r, c+1)
			processNeighbor(r, c-1)
		}
		ans++
	}
	return -1
}
