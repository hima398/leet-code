package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckPossibility(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, true, checkPossibility([]int{4, 2, 3}))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, false, checkPossibility([]int{4, 2, 1}))

	})
}
