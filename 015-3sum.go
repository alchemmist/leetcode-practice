//go:build ignore

package main

import "fmt"

func mergeSort(nums []int) []int {
    if len(nums) <= 1 {
        return nums
    }

    left := mergeSort(nums[:len(nums)/2])
    right := mergeSort(nums[len(nums)/2:])

    merged := []int{}
    i := 0
    j := 0
    for i < len(left) && j < len(right) {
        if left[i] <= right[j] {
            merged = append(merged, left[i])
            i++
        } else {
            merged = append(merged, right[j])
            j++
        }
    }
    merged = append(merged, left[i:]...)
    merged = append(merged, right[j:]...)

    return merged
}


func threeSum(nums []int) [][]int {
    nums = mergeSort(nums)
    var result [][]int
	n := len(nums)

    for i := 0; i < n-2; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        if nums[i] > 0 {
            break
        }

        l := i + 1
        r := n - 1

        for l < r {
            s := nums[i] + nums[l] + nums[r]
            if s < 0 {
                l++
            } else if s > 0 {
                r--
            } else {
                result = append(result, []int{nums[i], nums[l], nums[r]})

                for l < r && nums[l] == nums[l+1] {
                    l++
                }
                for l < r && nums[r] == nums[r-1] {
                    r--
                }

                l++
                r--
            }
        }
    }
    return result
}

func main() {
    fmt.Println(threeSum([]int{-1,0,1,2,-1,-4}))
}
