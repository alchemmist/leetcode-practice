//go:build ignore

package main

import (
	"fmt"
	"math"
)

func myAtoi(s string) int {
	negative := 1
	var result int = 0
	start := true
	for i := 0; i < len(s); i++ {
		if start && s[i] == ' ' {
			continue
		}
		if start && s[i] == '-' {
			negative = -1
			start = false
			continue
		} else if start && s[i] == '+' {
			negative = 1
			start = false
			continue
		}
		if !(s[i] >= '0' && s[i] <= '9') {
			break
		} else {
			start = false
		}
		dInt := int(s[i] - '0')
		if math.MinInt32 >= 10*result*negative-dInt {
			return math.MinInt32
		}
		if math.MaxInt32 <= 10*result+dInt*negative {
			return math.MaxInt32
		}
		result *= 10
		result += dInt
	}

	result *= negative
	return result
}

func main() {
	tests := []struct {
		in   string
		want int
	}{
		{"0", 0},
		{"42", 42},
		{"   42", 42},
		{"-42", -42},
		{"+42", 42},
		{"00123", 123},
		{"-00123", -123},
		{"1337c0d3", 1337},
		{"0-1", 0},
		{"-91283472332", -2147483648},
		{"   +0 123", 0},
		{"2147483646", 2147483646},
		{"-2147483648", -2147483648},
		{"-2147483647", -2147483647},
	}

	for _, tt := range tests {
		got := myAtoi(tt.in)
		if got != tt.want {
			fmt.Printf("myAtoi(%q) = %d, want %d\n", tt.in, got, tt.want)
		} else {
			fmt.Println("✓")
		}
	}
}
