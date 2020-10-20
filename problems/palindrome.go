package problems

func isPalindrome(x int) bool {
	if x < 0 || x%10 == 0 && x != 0 {
		return false
	}
	var y = 0
	for {
		y = y*10 + x%10
		x /= 10
		if y >= x {
			break
		}
	}
	return x == y || x == y/10
}
