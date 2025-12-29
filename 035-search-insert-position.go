//go:build ignore

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		i := left + (right-left)/2
		if nums[i] == target {
			return i
		}
		if target < nums[i] {
			right = i
		} else {
			left = i + 1
		}
	}
	return left
}

func main() {
	data := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(data, 6))
}
