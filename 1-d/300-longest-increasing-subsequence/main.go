package main

import "sort"

func main() {
	lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
}

func lengthOfLIS(nums []int) int {
	var piles [][]int
	for _, num := range nums {
		// Base case, we start the game with a new pile containing the first "card"
		if len(piles) == 0 {
			piles = append(piles, []int{num})
			continue
		}
		// This is the binary search portion
		// I'm not writing my own binary search, I'll just let Go do it
		// We want to find the left-most pile where the top card is
		// greater than or equal to the new card.
		j := sort.Search(len(piles), func(k int) bool {
			return piles[k][len(piles[k])-1] >= num
		})
		// If we found an appropriate pile, we add to it.
		// Otherwise, we start a new pile.
		if j < len(piles) {
			piles[j] = append(piles[j], num)
		} else {
			piles = append(piles, []int{num})
		}
	}
	// We don't actually care about any of the subsequences or merging the piles.
	// This is simply an early exit from a full-on Patience Sort.
	// The longest subsequence(s) are going to contain cards from every pile
	// so we just need the total number of piles and we have our answer.
	return len(piles)
}

// func lengthOfLIS(nums []int) int {
// 	dp := make([]int, len(nums))
// 	var result int

// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] < nums[j] && dp[i]+1 > dp[j] {
// 				dp[j] = dp[i] + 1
// 			}
// 		}
// 		result = max(result, dp[i])
// 	}

// 	return result + 1
// }

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
