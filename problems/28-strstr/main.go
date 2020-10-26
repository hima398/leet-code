package main

import "regexp"

// TODO: implement without regexp
func strStr(haystack string, needle string) int {
	r, _ := regexp.Compile(needle)
	loc := r.FindIndex([]byte(haystack))
	if len(loc) != 0 {
		return loc[0]
	} else {
		return -1
	}
}
