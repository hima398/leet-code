package main

func Question(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSubArray(nums []int) int {
	sum, max := nums[0], nums[0]
	for _, v := range nums[1:] {
		if v > v+sum {
			sum = v
		} else {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
