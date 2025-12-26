//go:build ignore

package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for i < len(nums) && j < len(nums) {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}
		i++
	}
	return j
}

func main() {
	data := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(data, 2))
	fmt.Println(data)
}
