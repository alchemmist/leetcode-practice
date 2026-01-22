//go:build ignore

package main

import "fmt"

func moveZeroes(nums []int)  {
    write := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != 0 {
            nums[write] = nums[i]
            write++
        }
    }
    for write < len(nums) {
        nums[write] = 0
        write++
    }
}

func main() {
    a := []int{0,0,1,3,12}
    moveZeroes(a)
    fmt.Println(a)
}
