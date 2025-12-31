//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0
	result := 0
	for l < len(s) && r < len(s) {
		if strings.Contains(s[l:r], string(s[r])) {
			result = max(r-l, result)
			l++
		} else {
			r++
		}
	}
	result = max(r-l, result)

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("au") == 2)
	fmt.Println(lengthOfLongestSubstring("aab") == 2)
}
