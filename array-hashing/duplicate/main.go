package main

func main() {
	containsDuplicate([]int{})
}

func containsDuplicate(nums []int) bool {
	n := make(map[int]struct{}, len(nums))

	for _, v := range nums {
		if _, ok := n[v]; ok {
			return true
		}
		n[v] = struct{}{}
	}

	return false
}
