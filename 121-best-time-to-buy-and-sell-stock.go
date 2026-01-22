//go:build ignore

package main

import "fmt"

func maxProfit(prices []int) int {
    mini := prices[0]
    maxi := 0
    for _, p := range prices {
        if p < mini {
            mini = p
        }
        if p - mini > maxi {
            maxi = p - mini
        }
    }
    return maxi
}

func main() {
    fmt.Println(maxProfit([]int{7,1,5,3,6,4}))
}
