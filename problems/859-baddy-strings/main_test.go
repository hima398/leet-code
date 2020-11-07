package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuddyStrings(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, true, buddyStrings("ab", "ba"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, false, buddyStrings("ab", "ab"))

	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, true, buddyStrings("aa", "aa"))

	})
	t.Run("Example 4", func(t *testing.T) {
		assert.Equal(t, true, buddyStrings("aaaaaaabc", "aaaaaaacb"))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.Equal(t, false, buddyStrings("", "aa"))
	})
	t.Run("Example 6", func(t *testing.T) {
		assert.Equal(t, false, buddyStrings("", ""))
	})
	t.Run("Example 7", func(t *testing.T) {
		assert.Equal(t, false, buddyStrings("a", "a"))
	})
	t.Run("Example 8", func(t *testing.T) {
		assert.Equal(t, false, buddyStrings("abcaa", "abcbb"))
	})
}
