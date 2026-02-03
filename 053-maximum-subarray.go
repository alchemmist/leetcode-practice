//go:build ignore

package main

import "math"

func maxSubArray(nums []int) int {
	curSum, maxSum := 0, math.Inf(-1)
	for _, n := range nums {
		curSum += n
		maxSum = max(maxSum, float64(curSum))

		if curSum < 0 {
			curSum = 0
		}
	}
	return int(maxSum)
}
