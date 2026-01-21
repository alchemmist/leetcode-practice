//go:build ignore

package main

import "fmt"

func containsDuplicate(nums []int) bool {
    haveSeen := make(map[int]struct{}, len(nums))
    for _, n := range nums {
        if _, ok := haveSeen[n]; ok {
            return true
        }
        haveSeen[n] = struct{}{}
    }
    return false
}

func main() {
    fmt.Println(
        containsDuplicate([]int{1, 2, 3, 1}),
    )
}
