//go:build ignore

package main

import "fmt"

func maxArea(height []int) int {
	i := 0
	j := len(height) - 1
	maxArea := 0
	for i < j {
		area := min(height[i], height[j]) * (j - i)
		if area > maxArea {
			maxArea = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
}
