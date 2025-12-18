//go:build ignore

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	for i_, i := range nums {

		for j_, j := range nums {
			if i_ != j_ && i + j == target {
				return []int{i_, j_}
			}
		}
	}
	return []int{0, 0}
}

func main() {
	var res []int = twoSum([]int{2,7,11,15}, 9);
	fmt.Printf("%d %d", res[0], res[1])
}
