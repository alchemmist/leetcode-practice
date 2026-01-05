//go:build ignore

package main

import "fmt"

func longestPalindrome(s string) string {
	odd := []rune{'|'}
	for _, chr := range s {
		odd = append(odd, chr, '|')
	}

	bestL, bestR := 0, 0

	for i := 0; i < len(odd); i++ {
		for shift := 0; i-shift >= 0 && i+shift < len(odd); shift++ {
			if odd[i-shift] != odd[i+shift] {
				break
			}
			if (i + shift - (i - shift)) > (bestR - bestL) {
				bestL = i - shift
				bestR = i + shift
			}
		}
	}
	res := make([]rune, 0, bestR-bestL+1)
	for _, r := range odd[bestL : bestR+1] {
		if r != '|' {
			res = append(res, r)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(longestPalindrome("бабад"))
}
