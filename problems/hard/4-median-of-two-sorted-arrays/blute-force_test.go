package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArrays(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		assert.Equal(t, 2.0, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("Example 2", func(t *testing.T) {
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		assert.Equal(t, 2.5, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("Example 3", func(t *testing.T) {
		nums1 := []int{0, 0}
		nums2 := []int{0}
		assert.Equal(t, 0.0, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("Example 4", func(t *testing.T) {
		nums1 := []int{}
		nums2 := []int{1}
		assert.Equal(t, 1.0, findMedianSortedArrays(nums1, nums2))
	})
	t.Run("Example 5", func(t *testing.T) {
		nums1 := []int{2}
		nums2 := []int{}
		assert.Equal(t, 2.0, findMedianSortedArrays(nums1, nums2))
	})
}

func BenchmarkFindMedianSortedArrays(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nums1 := []int{1, 3}
		nums2 := []int{2}
		findMedianSortedArrays(nums1, nums2)
	}
}
