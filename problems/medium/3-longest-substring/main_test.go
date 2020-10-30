package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLongestSubstring("bbbbb"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("pwwkew"))
	})
	t.Run("Example 4", func(t *testing.T) {
		assert.Equal(t, 0, lengthOfLongestSubstring(""))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.Equal(t, 3, lengthOfLongestSubstring("dvdf"))
	})
	t.Run("Example 6", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLongestSubstring("a"))
	})
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring("abcabcbb")
	}
}
