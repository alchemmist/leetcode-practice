//go:build ignore

package main

import (
	"fmt"
	"math"
)

func digitToInt(d rune) int {
	switch d {
	case '0':
		return 0
	case '1':
		return 1
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	}
	return -1
}

func myAtoi(s string) int {
	negative := 1
	var result int = 0
	start := true
	for _, d := range s {
		if start && d == ' ' {
			continue
		}
		if start && d == '-' {
			negative = -1
			start = false
			continue
		} else if start && d == '+' {
			negative = 1
			start = false
			continue
		}
		if digitToInt(d) == -1 {
			break
		} else {
			start = false
		}
		if math.MinInt32 >= 10*result*negative-digitToInt(d) {
			return math.MinInt32
		}
		if math.MaxInt32 <= 10*result+digitToInt(d)*negative {
			return math.MaxInt32
		}
		result *= 10
		result += digitToInt(d)
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
