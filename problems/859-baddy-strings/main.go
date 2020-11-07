package main

func buddyStrings(A string, B string) bool {
	na := len(A)
	nb := len(B)
	if na != nb {
		return false
	}
	var idx []int
	haveSameChar := false
	m := make(map[byte]int)
	for i := 0; i < na; i++ {
		if A[i] != B[i] {
			idx = append(idx, i)

			if len(idx) > 2 {
				return false
			}
		}
		haveSameChar = haveSameChar || m[A[i]] >= 1
		m[A[i]]++
	}
	if len(idx) < 2 {
		if len(idx) == 1 {
			return false
		}
		return haveSameChar
	}
	return A[idx[0]] == B[idx[1]] && A[idx[1]] == B[idx[0]]
}
