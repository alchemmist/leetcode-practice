//go:build ignore

package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	number := strconv.Itoa(x)
	l, r := 0, len(number) - 1
	for l < r {
		if number[l] != number[r] {
			return false
		}
		l += 1
		r -= 1
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(121))
}
