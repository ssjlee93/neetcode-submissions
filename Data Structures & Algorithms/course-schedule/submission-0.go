func canFinish(numCourses int, prerequisites [][]int) bool {
    // seeing possible route
    // DFS since we need one line direction.
    // prerequisites = adjacency list.
    // form a graph and see if a path exists.
    // prepend to each neighbors list.
 graph := make([][]int, numCourses)

    for _, p := range prerequisites {
        course, prerequisite := p[0], p[1]
        graph[course] = append(graph[course], prerequisite)
    }

    // 0 = unvisited, 1 = currently visiting, 2 = completed
    state := make([]int, numCourses)

    var dfs func(int) bool
    dfs = func(course int) bool {
        if state[course] == 1 {
            // Found a node already in the current DFS path.
            return false
        }

        if state[course] == 2 {
            // Already checked and found to be safe.
            return true
        }

        state[course] = 1

        for _, prerequisite := range graph[course] {
            if !dfs(prerequisite) {
                return false
            }
        }

        state[course] = 2
        return true
    }

    for course := 0; course < numCourses; course++ {
        if !dfs(course) {
            return false
        }
    }

    return true
}
