//go:build ignore

package main

func trailingZeroes(n int) int {
	result := 0
	i := 5
	for n/i > 0 {
		result += n / i
		i *= 5
	}
	return result
}
