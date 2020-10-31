package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		assert.Equal(t, 2, searchInsert(nums, 5))
	})
	t.Run("Example 2", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		assert.Equal(t, 1, searchInsert(nums, 2))
	})
	t.Run("Example 3", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		assert.Equal(t, 4, searchInsert(nums, 7))
	})
	t.Run("Example 4", func(t *testing.T) {
		nums := []int{1, 3, 5, 6}
		assert.Equal(t, 0, searchInsert(nums, 0))
	})
	t.Run("Example 5", func(t *testing.T) {
		nums := []int{1}
		assert.Equal(t, 0, searchInsert(nums, 0))
	})
}
