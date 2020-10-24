package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
		assert.True(t, isValid("{[]}"))
	})
	t.Run("Example 6", func(t *testing.T) {
		assert.False(t, isValid("(}[[[[[[[["))
	})
	t.Run("Example 7", func(t *testing.T) {
		assert.False(t, isValid(")"))
	})
	t.Run("Example 8", func(t *testing.T) {
		assert.False(t, isValid("}"))
	})
	t.Run("Example 9", func(t *testing.T) {
		assert.False(t, isValid("]"))
	})
}
