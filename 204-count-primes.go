//go:build ignore

package main

import (
	"fmt"
)

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}

	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i * i; j < n; j += i {
				primes[j] = false
			}
		}
	}

	count := 0
	for i := range n {
		if primes[i] {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}
