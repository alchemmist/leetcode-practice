//go:build ignore

package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	leftFinded, rightFinded := false, false
	for l <= r {
		i := (l + r) / 2
		if nums[i] == target {
			if nums[l] != target {
				l++
			} else {
				leftFinded = true
			}
			if nums[r] != target {
				r--
			} else {
				rightFinded = true
			}
			if leftFinded && rightFinded {
				return []int{l, r}
			}
		} else {
			if nums[i] < target {
				l = i + 1
			} else {
				r = i - 1
			}
		}
	}
	return []int{-1, -1}
}

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}
