//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	mem := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		if strings.Contains("({[", string(s[i])) {
			stack = append(stack, s[i])
			continue
		}
		if len(stack) == 0 {
			return false
		}
		if mem[s[i]] == stack[len(stack) - 1] {
			stack = stack[:len(stack) - 1]
		} else {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
}
