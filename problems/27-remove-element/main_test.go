package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Contains(expected []int, v int) bool {
	for i := 0; i < len(expected); i++ {
		if v == expected[i] {
			return true
		}
	}
	return false
}

func IsEqual(expected, actual []int, len int) bool {
	result := true
	for i := 0; i < len; i++ {
		if !Contains(expected, actual[i]) {
			return false
		}
	}
	return result
}

func TestRemoveElement(t *testing.T) {

	t.Run("Example 1", func(t *testing.T) {
		expected := []int{2, 2}
		actual := []int{3, 2, 2, 3}
		len := removeElement(actual, 3)
		assert.Equal(t, 2, len)
		assert.True(t, IsEqual(expected, actual, len))
	})
}
