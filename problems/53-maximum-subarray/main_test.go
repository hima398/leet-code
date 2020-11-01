package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 6, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, 1, maxSubArray([]int{1}))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 0, maxSubArray([]int{0}))
	})
	t.Run("Example 4", func(t *testing.T) {
		assert.Equal(t, -1, maxSubArray([]int{-1}))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.Equal(t, 2147483647, maxSubArray([]int{1, -2, 3, -4, 5, -6, 2147483647}))
	})
}
