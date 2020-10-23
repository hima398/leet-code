package main

import (
	"fmt"
	"sort"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	sort.Strings(strs)
	fmt.Println(strs)

	first := strs[0]
	last := strs[len(strs)-1]
	var result string
	for i := 0; i < len(first) && i < len(last); i++ {
		if first[i] != last[i] {
			break
		}
		result += string(first[i])
	}
	return result
}

/*
	if len(strs) < 2 {
		return ""
	}
	result := ""
	s := strs[0]
	for i := 0; i < len(s); i++ {
		for j := 1; j < len(str); j++ {
			if s[i] != str[j][i] {

			}

		}
	}

	return result
*/
