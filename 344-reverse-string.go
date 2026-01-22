//go:build ignore

package main

import "fmt"

func reverseString(s []byte)  {
    l := 0
    r := len(s) - 1
    for range len(s) / 2 {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
}

func main() {
    a := []byte{'h', 'e', 'l', 'l', 'o'}
    reverseString(a)
    fmt.Printf("%s\n", a)
}
