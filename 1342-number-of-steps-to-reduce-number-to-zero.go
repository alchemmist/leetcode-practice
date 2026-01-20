// go:build ignore

package main

import (
    "fmt"
)

func numberOfSteps(num int) int {
    if num == 0 {
        return 0
    }
    if num % 2 != 0 {
        return 1 + numberOfSteps(num - 1)
    } else {
        return 1 + numberOfSteps(num / 2)
    }
}

func main() {
    fmt.Println(numberOfSteps(123))
}
