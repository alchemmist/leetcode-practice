//go:build ignore

package main

import "fmt"

func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}
	a := len(str1)
	b := len(str2)
	for a > 0 && b > 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return str1[:max(a, b)]
}

func main() {
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
}

