//go:build ignore

package main

import "fmt"

func removeElement(nums []int, val int) int {
	removed := 0
	i := 0
	for {
		if i >= len(nums) {
			break
		}
		if nums[i] == val {
			removed += 1
			for j := 1; j < len(nums)-i; j++ {
				nums[i+j-1] = nums[i+j]
			}
			nums[len(nums)-1] = val + 1
		} else {
			i++
		}
	}
	return len(nums) - removed
}

func main() {
	data := []int{4, 4, 0, 1, 0, 2}
	fmt.Println(removeElement(data, 0))
	fmt.Println(data)
}
