//go:build ignore

package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var cache  = make(map[int]int)
	for i, n := range nums {
		if j, ok := cache[n]; ok {
			return []int{j, i}
		}
		cache[target - n] = i
	}
	return nil
}

func main() {
	var res []int = twoSum([]int{2,7,11,15}, 9);
	fmt.Printf("%d %d", res[0], res[1])
}
