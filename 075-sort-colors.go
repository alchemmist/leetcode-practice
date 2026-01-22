//go:build ignore

package main

import "fmt"

func sortColors(nums []int)  {
	l, i, r := 0, 0, len(nums)-1
	for i <= r {
		if nums[i] == 0 {
			nums[l], nums[i] = nums[i], nums[l]
			l++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
}

func main() {
	a := []int{2,0,2,1,1,0}
	sortColors(a)
	fmt.Println(a)
}
