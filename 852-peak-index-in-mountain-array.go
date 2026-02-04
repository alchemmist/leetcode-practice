//go:build ignore

package main

import (
	"fmt"
)

func peakIndexInMountainArray(arr []int) int {
	l, r := 0, len(arr) - 1
	for l < r {
		i := (r + l) / 2
		if arr[i] > arr[i + 1] {
			r = i
		} else {
			l = i + 1
		}
	}
	return l 
}

func main() {
	fmt.Println(peakIndexInMountainArray([]int{0, 10, 5, 2}))
}
