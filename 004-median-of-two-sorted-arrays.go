//go:build ignore

package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := []int{}
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	nums = append(nums, nums1[i:]...)
	nums = append(nums, nums2[j:]...)
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2]+nums[len(nums)/2-1]) / float64(2)
	} else {
		return float64(nums[len(nums)/2])
	}
}

func main() {
	a := []int{1, 3}
	b := []int{2}
	fmt.Println(findMedianSortedArrays(a, b))

	a = []int{1, 2}
	b = []int{3, 4}
	fmt.Println(findMedianSortedArrays(a, b))
}
