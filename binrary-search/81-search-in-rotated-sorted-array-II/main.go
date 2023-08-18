package main

func main() {
	search([]int{1, 0, 1, 1, 1}, 0)
}

func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}

		if nums[l] == nums[mid] {
			l++
			continue
		}

		if nums[l] <= nums[mid] {
			if nums[l] <= target && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return false
}
