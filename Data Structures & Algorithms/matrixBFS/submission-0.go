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
			neighbors := []Pair{ {1, 0}, {-1, 0}, {0, 1}, {0, -1} }
			for _, p := range neighbors {
				dr, dc := p.Row, p.Col
				rd, cd := r+dr, c+dc
				point := Pair{rd, cd}

				// boundary
				if rd < 0 || cd < 0 || rd == ROWS || cd == COLS ||
				visited[point] ||
				grid[rd][cd] == 1 {
					continue
				}

				// process our queue
				q = append(q, point)
				visited[point] = true
			}
			
		}
		ans++
	}
	return -1
}
