//go:build ignore

package main

import "fmt"

func romanToInt(s string) int {
	roman_numbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var number int = 0
	for i := len(s) - 1; i >= 0; i-- {
		r := roman_numbers[s[i]]
		if i == 0 {
			number += r
			break
		}

		l := roman_numbers[s[i-1]]
		if r > l {
			number += r - l
			i--
		} else {
			number += r
		}
	}

	return number
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
