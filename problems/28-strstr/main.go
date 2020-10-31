package main

func strStr(haystack string, needle string) int {
	nlen := len(needle)
	if nlen == 0 {
		return 0
	}
	hlen := len(haystack)
	if hlen == 0 {
		return -1
	}

	c := needle[0]
	for i := 0; i < hlen-nlen+1; i++ {
		if haystack[i] == c {
			subStr := haystack[i : i+nlen]
			if subStr == needle {
				return i
			}
		}
	}
	return -1
}

// TODO: implement without regexp
/*
func strStr(haystack string, needle string) int {
	r, _ := regexp.Compile(needle)
	loc := r.FindIndex([]byte(haystack))
	if len(loc) != 0 {
		return loc[0]
	} else {
		return -1
	}
}
*/
