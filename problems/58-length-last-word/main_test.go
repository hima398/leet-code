package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		assert.Equal(t, 5, lengthOfLastWord("Hello World"))
	})
	t.Run("Example2", func(t *testing.T) {
		assert.Equal(t, 1, lengthOfLastWord("a "))
	})
}
