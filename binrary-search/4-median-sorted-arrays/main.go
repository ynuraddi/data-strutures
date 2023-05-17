package main

func main() {
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var num1, num2 int

	for i, j, count := 0, 0, 0; count < (len(nums1)+len(nums2))/2; {
		if i < len(nums1) && (j == len(nums2) || nums1[i] <= nums2[j]) {
			num1, num2 = num2, nums1[i]
			i++
			count++
		}

		if j < len(nums2) && (i == len(nums1) || nums2[j] <= nums1[i]) {
			num1, num2 = num2, nums2[j]
			j++
			count++
		}

	}

	return float64(num1+num2) / 2
}
