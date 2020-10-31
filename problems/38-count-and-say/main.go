package main

import (
	"strconv"
)

func countAndSay(n int) string {
	if n <= 1 {
		return "1"
	}
	// n > 1
	s := countAndSay(n - 1)
	sl := len(s)
	c := s[0]
	//count := 1
	count := 0
	substr := ""
	for i := 0; i < sl; i++ {
		if i < sl && c == s[i] {
			count++
		} else {
			substr += strconv.Itoa(count) + string(c)
			c = s[i]
			count = 1
		}
	}
	substr += strconv.Itoa(count) + string(c)
	return substr
}
