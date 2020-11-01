package main

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := append(nums1, nums2...)
	sort.Ints(nums)
	n := len(nums)
	if n%2 == 1 {
		// odd
		return float64(nums[n/2])
	} else {
		// even
		return float64(nums[n/2]+nums[n/2-1]) / 2
	}
}
