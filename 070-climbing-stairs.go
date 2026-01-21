//go:build inore

package main

import "fmt"

func climbStairs(n int) int {
    if n == 1 {
        return 1
    }
    a, b := 0, 1
    for range n {
        a, b = b, a+b
    }
    return b
}

func main() {
    fmt.Println(climbStairs(37))
}
