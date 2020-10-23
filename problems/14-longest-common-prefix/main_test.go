package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		assert.Equal(t, "fl", longestCommonPrefix(strs))
	})
	t.Run("Example2", func(t *testing.T) {
		strs := []string{"dog", "racecar", "car"}
		assert.Equal(t, "", longestCommonPrefix(strs))
	})
	t.Run("Empty strs", func(t *testing.T) {
		strs := []string{}
		assert.Equal(t, "", longestCommonPrefix(strs))
	})
	t.Run("One strs", func(t *testing.T) {
		strs := []string{"strs"}
		assert.Equal(t, "strs", longestCommonPrefix(strs))
	})
}
