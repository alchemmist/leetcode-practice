//go:build ignore

package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if len(s) <= 2 {
		return s
	}
	if numRows == 1 {
		return s
	}
	zigzag := []string{}
	for range numRows {
		zigzag = append(zigzag, "")
	}
	cycle := (numRows - 2) + numRows
	for i := 0; i < len(s); i++ {
		if (i % cycle) < numRows {
			zigzag[i%cycle] += string(s[i])
		} else {
			zigzag[cycle-(i%cycle)] += string(s[i])
		}
	}
	return strings.Join(zigzag, "")
}

func main() {
	fmt.Println(convert("PAYPALISHIRING", 4) == "PINALSIGYAHRPI")
	fmt.Println(convert("PAYPALISHIRING", 3) == "PAHNAPLSIIGYIR")
	fmt.Println(convert("A", 1) == "A")
	fmt.Println(convert("ABC", 1) == "ABC")
}
