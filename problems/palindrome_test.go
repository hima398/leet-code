package problems

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		assert.True(t, isPalindrome(121))
	})
	t.Run("Example2", func(t *testing.T) {
		assert.False(t, isPalindrome(-121))
	})
	t.Run("Example3", func(t *testing.T) {
		assert.False(t, isPalindrome(10))
	})
	t.Run("Example4", func(t *testing.T) {
		assert.False(t, isPalindrome(-101))
	})
	t.Run("偶数桁", func(t *testing.T) {
		assert.True(t, isPalindrome(1221))
	})
	t.Run("int32の最大値と同じ桁", func(t *testing.T) {
		assert.True(t, isPalindrome(1234554321))
	})
}
