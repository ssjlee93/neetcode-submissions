func canFinish(numCourses int, prerequisites [][]int) bool {
    graph := make([][]int, numCourses)
    indegree := make([]int, numCourses)

    for _, p := range prerequisites {
        course, prerequisite := p[0], p[1]

        // prerequisite -> course
        graph[prerequisite] = append(graph[prerequisite], course)
        indegree[course]++
    }

    queue := []int{}

    for course := 0; course < numCourses; course++ {
        if indegree[course] == 0 {
            queue = append(queue, course)
        }
    }

    completed := 0

    for len(queue) > 0 {
        course := queue[0]
        queue = queue[1:]
        completed++

        for _, nextCourse := range graph[course] {
            indegree[nextCourse]--

            if indegree[nextCourse] == 0 {
                queue = append(queue, nextCourse)
            }
        }
    }

    return completed == numCourses
}
