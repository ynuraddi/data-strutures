package main

func main() {
}

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if index, ok := hash[target-nums[i]]; ok {
			return []int{index, i}
		} else {
			hash[nums[i]] = i
		}
	}

	return []int{}
}
