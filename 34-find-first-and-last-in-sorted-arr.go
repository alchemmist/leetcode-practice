//go:build ignore

package main

import (
	"fmt"
)

func lowerBound(nums []int, target int) int {
	if len(nums) == 1 {
		return 0
	}
	l, r := 0, len(nums)
	for l < r {
		i := (l + r) / 2
		if nums[i] < target {
			l = i + 1
		} else {
			r = i
		}
	}
	return l
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		i := (l + r) / 2
		if nums[i] <= target {
			l = i + 1
		} else {
			r = i
		}
	}
	return l
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	first := lowerBound(nums, target)

	if first == len(nums) || nums[first] != target {
		return []int{-1, -1}
	}

	last := upperBound(nums, target) - 1
	return []int{first, last}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
