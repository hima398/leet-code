package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	var substr string
	var j int
	maxLen := 0
	for i := 0; j < n; i++ {
		substr = string(s[i])
		for j = i + 1; j < n; j++ {
			if strings.Contains(substr, string(s[j])) {
				break
			} else {
				substr = s[i : j+1]
			}
		}
		if len(substr) > maxLen {
			maxLen = len(substr)
		}
	}
	return maxLen
}
