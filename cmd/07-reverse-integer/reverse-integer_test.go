package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	t.Run("example1", func(t *testing.T) {
		assert.Equal(t, 321, reverse(123))
	})
	t.Run("example2", func(t *testing.T) {
		assert.Equal(t, -321, reverse(-123))
	})
	t.Run("example3", func(t *testing.T) {
		assert.Equal(t, 21, reverse(120))
	})
	t.Run("example4", func(t *testing.T) {
		assert.Equal(t, 0, reverse(0))
	})
	t.Run("reverse = max int(2 ^ 31 - 1)", func(t *testing.T) {
		// 2147483647: Max int32
		assert.Equal(t, 2147483647, reverse(7463847412))
	})
	t.Run("reverse = max int(2 ^ 31)", func(t *testing.T) {
		// Reverse = 2147483647 + 1
		assert.Equal(t, 0, reverse(8463847412))
	})
	t.Run("max int(-2 ^ 31)", func(t *testing.T) {
		// Reverse = -2147483648
		assert.Equal(t, -2147483648, reverse(-8463847412))
	})
	t.Run("max int(-2 ^ 31)", func(t *testing.T) {
		// Reverse = -2147483648 - 1
		assert.Equal(t, 0, reverse(-9463847412))
	})
}
