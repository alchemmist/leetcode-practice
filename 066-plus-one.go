//go:build ignore

package main

import (
	"fmt"
	"slices"
)

func plusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; i >= 0; i-- {
		tmp := digits[i] + remainder
		digits[i] = tmp % 10
		remainder = tmp / 10
	}
	if remainder != 0 {
		digits = slices.Insert(digits, 0, remainder)
	}
	return digits
}

func main() {
	n := []int{9}
	fmt.Println(plusOne(n))
}
