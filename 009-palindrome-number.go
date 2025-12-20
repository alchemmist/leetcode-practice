//go:build ignore

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	orig := x
	rev := 0
	for x > 0 {
		rev = rev*10 + x%10
		x /= 10
	}
	return rev == orig
}

func main() {
	fmt.Println(isPalindrome(100))
}
