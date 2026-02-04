//go:build ignore

package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		i := (l + r) / 2
		if nums[i] == target {
			return i
		}
		if nums[l] <= nums[i] {
			if nums[l] <= target && target <= nums[i] {
				r = i - 1
			} else {
				l = i + 1
			}
		} else {
			if nums[i] <= target && target <= nums[r] {
				l = i + 1
			} else {
				r = i - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 7))
}
