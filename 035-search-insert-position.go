//go:build ignore

package main

import "fmt"

func searchInsert(nums []int, target int) int {
	res := len(nums)
	for i, n := range nums {
		if n == target {
			return i
		}
		if target < n {
			res = min(i, res)
		}
	}
	return res
}

func main() {
	data := []int{1, 3, 5, 6}
	fmt.Println(searchInsert(data, 5))
}
