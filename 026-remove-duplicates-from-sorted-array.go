//go:build ignore

package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	i := 1
	k := 1
	for {
		if i == len(nums) - 1 || nums[i] == nums[len(nums) - 1] {
			break
		}
		if nums[i] == nums[i - 1] {
			for shift := 1; i + shift < len(nums); shift++ {
				nums[i + shift - 1] = nums[i + shift]
			}
		} 
		if nums[i] != nums[i - 1] {
			i ++
			k ++
		}
	}
	if nums[i] != nums[i - 1] {
		i ++
		k ++
	}
	return k
}

func main() {
	var data []int = []int{1, 1, 2, 3, 3, 3, 3, 4}
	fmt.Println(removeDuplicates(data), data)
}
