//go:build ignore

package main

import "fmt"

func isBadVersion(version int) bool {
    return version >= 4
}

func firstBadVersion(n int) int {
    l := 0
    r := n
    for l <= r {
        i := l + (r - l) / 2
        if isBadVersion(i) {
            r = i - 1
        } else {
            l = i + 1
        }
    }
    return l
}

func main() {
    fmt.Println(firstBadVersion(5))
}
