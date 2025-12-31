//go:build ignore

package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	l := 0
	r := 0
	result := 0
	lastPos := make(map[byte]int)
	for l < len(s) && r < len(s) {
		p, ok := lastPos[s[r]]
		if ok && p >= l {
			l = p + 1
		}
		lastPos[s[r]] = r
		result = max(r-l+1, result)
		r++
	}

	return result
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb") == 3)
	fmt.Println(lengthOfLongestSubstring("bbb") == 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew") == 3)
	fmt.Println(lengthOfLongestSubstring("au") == 2)
	fmt.Println(lengthOfLongestSubstring("aab") == 2)
}
