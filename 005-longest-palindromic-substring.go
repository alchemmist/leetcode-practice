//go:build ignore

package main

import "fmt"

func longestPalindrome(s string) string {
	s_odd := "|"
	for _, chr := range s {
		s_odd += string(chr) + "|"
	}

	maxPal := ""
	for i := 0; i < len(s_odd); i++ {
		for shift := 0; i+shift < len(s_odd) && 0 <= i-shift; shift++ {
			if s_odd[i+shift] == s_odd[i-shift] {
				if len(s_odd[i-shift:i+shift+1]) > len(maxPal) {
					maxPal = s_odd[i-shift : i+shift+1]
				}
			} else {
				break
			}
		}
	}
	res := ""
	for _, letter := range maxPal {
		if letter != '|' {
			res = res + string(letter)
		}
	}
	return res
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
