package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMap(t *testing.T) {
	var pMap = map[byte]byte{
		40:  41,  // ()
		123: 125, // {}
		91:  93,  // []
	}
	assert.Equal(t, 41, pMap[40])
	assert.Equal(t, 125, pMap[123])
	assert.Equal(t, 93, pMap[91])

}

func TestIsValid(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.True(t, isValid("()"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.True(t, isValid("()[]{}"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.False(t, isValid("(]"))
	})
	t.Run("Example 4", func(t *testing.T) {
		assert.False(t, isValid("([)]"))
	})
	t.Run("Example 5", func(t *testing.T) {
		assert.False(t, isValid("{[]}"))
	})
}
