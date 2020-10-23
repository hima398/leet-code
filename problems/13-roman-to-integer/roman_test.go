package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRomanToInt(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		assert.Equal(t, 3, romanToInt("III"))
	})
	t.Run("Example2", func(t *testing.T) {
		assert.Equal(t, 4, romanToInt("IV"))
	})
	t.Run("Example3", func(t *testing.T) {
		assert.Equal(t, 9, romanToInt("IX"))
	})
	t.Run("Example4", func(t *testing.T) {
		assert.Equal(t, 58, romanToInt("LVIII"))
	})
	t.Run("Example5", func(t *testing.T) {
		assert.Equal(t, 1994, romanToInt("MCMXCIV"))
	})
	t.Run("Min Case", func(t *testing.T) {
		assert.Equal(t, 1, romanToInt("I"))
	})
	t.Run("Max Value Case", func(t *testing.T) {
		assert.Equal(t, 3999, romanToInt("MMMCMXCIX"))
	})
	t.Run("Max Length Case", func(t *testing.T) {
		assert.Equal(t, 3888, romanToInt("MMMDCCCLXXXVIII"))
	})
}
