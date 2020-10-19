package main

const MaxInt = 2147483647
const MinInt = -2147483648

func reverse(x int) int {
	reverse := 0
	for {
		if x == 0 {
			break
		}
		reverse = reverse*10 + x%10
		if reverse > MaxInt {
			return 0
		}
		if reverse < MinInt {
			return 0
		}
		x /= 10
	}
	return reverse
}
