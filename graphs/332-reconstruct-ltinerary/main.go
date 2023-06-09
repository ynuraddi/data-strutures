package main

import "sort"

func main() {
	findItinerary([][]string{
		{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"},
	})
}

func findItinerary(tickets [][]string) (res []string) {
	adjMap := make(map[string][]string)

	for _, t := range tickets {
		adjMap[t[0]] = append(adjMap[t[0]], t[1])
	}

	for t := range adjMap {
		sort.Strings(adjMap[t])
	}

	var dfs func(s string) bool
	dfs = func(s string) bool {
		if len(res) == len(tickets) {
			return true
		}

		if _, ok := adjMap[s]; !ok {
			return false
		}

		destinationList := adjMap[s]
		for i, dest := range destinationList {
			adjMap[s] = append(adjMap[s][:i], adjMap[s][i+1:]...)

			res = append(res, dest)
			if dfs(dest) {
				return true
			}
			adjMap[s] = append(adjMap[s][:i+1], adjMap[s][i:]...)
			adjMap[s][i] = dest

			res = res[:len(res)-1]
		}
		return false
	}
	dfs("JFK")
	return append([]string{"JFK"}, res...)
}
