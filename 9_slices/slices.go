package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// most used construct in go

func main() {
	// uninitialized slice is nil by default
	// var nums []int
	// fmt.Println(nums) // Prints []
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// var nums = make([]int, 0, 5)
	// fmt.Println(nums) // Prints [0 0]
	// capacity -> Maximum number of elements that can fit
	// fmt.Println(cap(nums))

	// var nums = make([]int, 0, 5)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 1)
	// nums = append(nums, 20)
	// nums = append(nums, 12)
	// fmt.Println(nums)      // [1 1 1 1 20 12]
	// fmt.Println(cap(nums)) // 10

	// nums := []int{}
	// nums = append(nums, 20)
	// fmt.Println(nums)

	// Copy function
	// var nums = make([]int, 0, 5)
	// nums = append(nums, 22)
	// var nums2 = make([]int, len(nums))
	// copy(nums2, nums)
	// fmt.Println(nums, nums2)

	// Slice Operator
	// var nums = []int{1, 2, 3}
	// fmt.Println(nums[0:2])

	// Slice
	var nums1 = []int{1, 2}
	var nums2 = []int{1, 2}
	fmt.Println(slices.Equal(nums1, nums2))

	var nums = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(nums)

}
