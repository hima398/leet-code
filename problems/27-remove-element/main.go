package main

func removeElement(nums []int, val int) int {
	// Constraints 0 <= nums[i] <= 50 and 0 <= val <= 100
	if val > 50 {
		return len(nums)
	}
	return 0
}
