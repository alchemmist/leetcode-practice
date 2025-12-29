//go:build ignore

package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length != 0 {
			return length
		}
		if s[i] != ' ' {
			length++
		}

	}
	return length
}

func main() {
	var s string = "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
