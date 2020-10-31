package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, "1", countAndSay(1))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, "11", countAndSay(2))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, "21", countAndSay(3))
	})
	t.Run("Example 4", func(t *testing.T) {
		assert.Equal(t, "1211", countAndSay(4))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.Equal(t, "111221", countAndSay(5))
	})
	t.Run("Example 6", func(t *testing.T) {
		assert.Equal(t, "312211", countAndSay(6))
	})
}
