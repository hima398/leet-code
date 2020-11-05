package main

func lengthOfLastWord(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	i := n - 1
	// 32 = space in ASCII Code
	for ; i >= 0 && s[i] == 32; i-- {

	}
	j := i
	for j >= 0 && s[j] != 32 {
		j--
	}
	return i - j
}
