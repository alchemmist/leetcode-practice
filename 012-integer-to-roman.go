//go:build ignore

package main

import "fmt"

type Roman struct {
	name  string
	value int
}

func intToRoman(num int) string {
	roman := []Roman{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	result := ""
	for num > 0 {
		for _, r := range roman {
			if r.value <= num {
				result += r.name
				num -= r.value
				break
			}
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(3749))
}
