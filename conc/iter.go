package main

import (
	"fmt"

	"github.com/sourcegraph/conc/iter"
)

// iter.ForEach operates elements in original slice
func tryIter() {
	nums := []int{0, 1, 2, 3, 4, 5, 6}

	handler := func(i *int) {
		*i++
	}

	iter.ForEach(nums, handler)
	fmt.Println(nums)
}

// iter.Map returns a new slice with elements processed by f
func concMap(input []int, f func(*int) int) []int {
	return iter.Map(input, f)
}

func tryIterMap() {
	nums := []int{0, 1, 2, 3, 4, 5, 6}
	handler := func(i *int) int {
		return *i + 1
	}

	nums2 := concMap(nums, handler)

	fmt.Println(nums2)
}
