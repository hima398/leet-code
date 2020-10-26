package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 2, strStr("Hello", "ll"))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, -1, strStr("aaaaa", "bba"))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 0, strStr("", ""))
	})
}
