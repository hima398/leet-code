package main

import "fmt"

func twoSum(nums []int, target int) []int {

	var m = make(map[int]int)
	len := len(nums)
	for i := 0; i < len; i++ {
		m[nums[i]] = i
	}
	fmt.Println(m)
	for i := 0; i < len; i++ {
		n := target - nums[i]
		if m[n] != 0 && i != m[n] {
			return []int{i, m[n]}
		}
	}
	return []int{-1, -1}
}

func main() {
	example1 := []int{2, 7, 11, 15}
	fmt.Println(twoSum(example1, 9))

	example2 := []int{3, 2, 4}
	fmt.Println(twoSum(example2, 6))

	example3 := []int{3, 3}
	fmt.Println(twoSum(example3, 6))

	example4 := []int{1, 3, 4, 2}
	fmt.Println(twoSum(example4, 6))
}
