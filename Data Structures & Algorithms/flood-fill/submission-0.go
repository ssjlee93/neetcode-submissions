type Pair struct {
    Row int
    Col int
}

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    // use bfs grid traversal
    // no need for back tracking since we only care about color and its state, not path
    ROWS, COLS := len(image), len(image[0])
    visited := make(map[Pair]bool)
    sColor := image[sr][sc]

    var dfs func(int, int, map[Pair]bool)
    dfs = func(r, c int, visit map[Pair]bool) {
        point := Pair{r, c}
        // base case : boundary
        if r < 0 || c < 0 || r == ROWS || c == COLS ||
        visited[point] ||
        sColor != image[r][c] {
            return
        }

        // mark visited
        visited[point] = true

        // cbange color
        image[r][c] = color

        // recurse neighbors
        dfs(r+1, c, visit)
        dfs(r-1, c, visit)
        dfs(r, c+1, visit)
        dfs(r, c-1, visit)
        return
    }

    dfs(sr, sc, visited)
    return image
}
