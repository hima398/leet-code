package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 0 {
		return 0
	}
	i := 0
	for j := 1; j < n; j++ {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
