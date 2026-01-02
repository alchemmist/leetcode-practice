//go:build ignore

package main

import "fmt"

func longestPalindrome(s string) string {
	s_odd := "|"
	for _, chr := range s {
		s_odd += string(chr) + "|"
	}
	for i := 0; i < len(s_odd); i++ {
		shift := 0
		for i + shift < len(s_odd) && 0 <= i - shift {
			if s_odd[i + shift] == s_odd[i - shift] {

			}
		}
	}
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
