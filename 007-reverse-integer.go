//go:build ignore

package main

import (
	"fmt"
	"math"
)

func pow10(n int) int {
	r := 1
	for range n {
		r *= 10
	}
	return r
}

func reverse(x int) int {
	absX := max(x, -x)
	result := 0
	for {
		if math.MaxInt32 - result < 9 * result + (absX % 10) {
			return 0
		}
		result = result*10 + (absX % 10)

		if absX/10 == 0 {
			break
		}
		absX /= 10
	}

	if x < 0 {
		result *= -1
	}
	return result
}

func main() {
	fmt.Println(reverse(123) == 321)
	fmt.Println(reverse(-123) == -321)
	fmt.Println(reverse(120) == 21)
	fmt.Println(reverse(1534236469) == 0)
}
