//go:build ignore

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	w := 0
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[w] {
			w++
			nums[w] = nums[r]
		}
	}
	return w + 1
}

func main() {
	var data []int = []int{1, 1, 2}
	fmt.Println(removeDuplicates(data), data)
}
