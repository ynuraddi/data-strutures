package main

func main() {
	canFinish(2, [][]int{{1, 0}, {0, 2}, {3, 2}})
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int)

	for _, prer := range prerequisites {
		graph[prer[1]] = append(graph[prer[1]], prer[0])
	}

	visited := make(map[int]struct{})

	var dfs func(course int) bool
	dfs = func(course int) bool {
		if _, ok := visited[course]; ok {
			return false
		}

		if len(graph) == 0 {
			return true
		}

		visited[course] = struct{}{}

		for _, pre := range graph[course] {
			if !dfs(pre) {
				return false
			}
		}
		delete(visited, course)

		graph[course] = []int{}

		return true
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i) {
			return false
		}
	}

	return true
}
