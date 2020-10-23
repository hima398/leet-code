package main

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	n := len(s)
	sum := 0
	for i := 0; i < n; i++ {
		if i+1 < n && roman[string(s[i])] < roman[string(s[i+1])] {
			sum -= roman[string(s[i])]
		} else {
			sum += roman[string(s[i])]
		}
	}
	return sum
}
