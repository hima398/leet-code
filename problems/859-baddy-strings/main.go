package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}
	if A == B {
		return false
	}
	return true
}
