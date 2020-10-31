package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	i := 0
	for ; i < n; i++ {
		if target <= nums[i] {
			break
		}
	}
	return i
}
