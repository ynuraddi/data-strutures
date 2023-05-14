package main

import (
	"fmt"
)

func main() {
	fmt.Println(minEatingSpeed([]int{312884470}, 312884469))
}

func minEatingSpeed(piles []int, h int) int {
	howLong := func(speed int) int {
		time := 0
		for _, pile := range piles {
			time += pile / speed
			if pile%speed != 0 {
				time++
			}
		}
		return time
	}
	max := 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	left, right, res := 1, max, max
	for left <= right {
		mid := left + (right-left)/2
		if howLong(mid) <= h {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
