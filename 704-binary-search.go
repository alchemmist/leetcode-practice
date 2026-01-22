//go:build ignore

package main

import "fmt"

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		i := l + (r - l) / 2
		if nums[i] < target {
			l = i + 1
		} else if nums[i] > target {
			r = i - 1
		} else {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1,0,3,5,9,12}, 2))
}
