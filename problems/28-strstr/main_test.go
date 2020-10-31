package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 2, strStr("Hello", "ll"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, -1, strStr("aaaaa", "bba"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 0, strStr("", ""))
	})
}

func BenchmarkStrStr(b *testing.B) {
	n := 50000
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = 97
	}
	haystack := string(buf)
	for i := 0; i < b.N; i++ {
		strStr(haystack, "b")
	}
}
