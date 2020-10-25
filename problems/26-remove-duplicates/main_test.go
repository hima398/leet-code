package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsEqual(expected, actual []int, n int) bool {
	result := true
	for i := 0; i < n; i++ {
		result = result && (expected[i] == actual[i])
	}
	return result
}

func TestRemoveDuplicates(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		expected := []int{1, 2}
		actual := []int{1, 1, 2}
		n := removeDuplicates(actual)
		assert.Equal(t, 2, n)
		assert.True(t, IsEqual(expected, actual, n))
	})
	t.Run("Example 2", func(t *testing.T) {
		expected := []int{0, 1, 2, 3, 4}
		actual := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		n := removeDuplicates(actual)
		assert.Equal(t, 5, n)
		assert.True(t, IsEqual(expected, actual, n))
	})
	t.Run("Example 3", func(t *testing.T) {
		expected := []int{}
		actual := []int{}
		n := removeDuplicates(actual)
		assert.Equal(t, 0, n)
		assert.True(t, IsEqual(expected, actual, n))
	})

}
