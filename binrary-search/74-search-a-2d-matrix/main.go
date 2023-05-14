package main

import "fmt"

func main() {
	fmt.Println(searchMatrix([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, -1))
}

func searchMatrix(matrix [][]int, target int) bool {
	l, r := 0, len(matrix)-1
	possiblyContainIndex := 0

	for l <= r {
		mid := (l + r) / 2
		if target > matrix[mid][0] {
			l = mid + 1
		} else if target < matrix[mid][0] {
			r = mid - 1
		} else {
			possiblyContainIndex = mid
			if matrix[mid][0] == target {
				return true
			}
		}
	}

	if r < 0 {
		return false
	}

	if possiblyContainIndex == 0 {
		possiblyContainIndex = r
	}

	l, r = 0, len(matrix[possiblyContainIndex])-1
	for l <= r {
		mid := (l + r) / 2
		if target > matrix[possiblyContainIndex][mid] {
			l = mid + 1
		} else if target < matrix[possiblyContainIndex][mid] {
			r = mid - 1
		} else {
			return true
		}
	}

	return false
}
