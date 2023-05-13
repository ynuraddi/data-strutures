package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{3}, []int{3}))
	fmt.Println(carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
	fmt.Println(carFleet(10, []int{8, 3, 7, 4, 6, 5}, []int{4, 4, 4, 4, 4, 4}))
}

func carFleet(target int, position []int, speed []int) int {
	// create an array of pairs to sort it
	// and not lose the relationship between speed and position
	pairs := make([][2]int, len(position))
	for i := range pairs {
		pairs[i][0] = position[i]
		pairs[i][1] = speed[i]
	}

	// sort by distance from the target
	// from the closest to the farthest
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] > pairs[j][0]
	})

	// the top stack will be the car at the tail end
	stack := make([]float64, 0, len(position))

	for i := range pairs {
		currentTime := float64(target-pairs[i][0]) / float64(pairs[i][1])

		// If it goes faster than the car with the tail, it means they are teaming up
		// and there is no need to fluff it in the text
		if len(stack) == 0 || currentTime > stack[len(stack)-1] {
			stack = append(stack, currentTime)
		}
	}

	return len(stack)
}
