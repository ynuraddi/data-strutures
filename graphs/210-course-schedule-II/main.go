package main

import "fmt"

func main() {
	fmt.Println(findOrder(4, [][]int{{1, 0}, {3, 0}, {2, 1}, {2, 3}}))
}

func findOrder(numCourses int, prerequisites [][]int) (res []int) {
	prereq := make([][]int, numCourses)
	for _, pre := range prerequisites {
		prereq[pre[0]] = append(prereq[pre[0]], pre[1])
	}

	visit, cycle := make([]bool, numCourses), make([]bool, numCourses)
	var dfs func(int) bool
	dfs = func(crs int) bool {
		if cycle[crs] {
			return false
		} else if visit[crs] {
			return true
		}

		cycle[crs] = true
		for _, pre := range prereq[crs] {
			if !dfs(pre) {
				return false
			}
		}
		cycle[crs] = false
		visit[crs] = true
		res = append(res, crs)
		return true
	}

	for c := 0; c < numCourses; c++ {
		if !dfs(c) {
			return []int{}
		}
	}
	return res
}
