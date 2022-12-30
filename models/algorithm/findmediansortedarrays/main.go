package main

import (
	"fmt"
)

func sortArr(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	n := nums[0]
	s := make([]int, 0)
	l := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= n {
			s = append(s, nums[i])
		} else {
			l = append(l, nums[i])
		}
	}

	ret := make([]int, 0)
	if len(s) > 0 {
		ret = append(ret, sortArr(s)...)
	}
	ret = append(ret, n)
	if len(l) > 0 {
		ret = append(ret, sortArr(l)...)
	}

	return ret
}

func mergeArr(nums1 []int, nums2 []int) []int {
	nums := make([]int, 0)
	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	ret := sortArr(nums)
	return ret
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums := mergeArr(nums1, nums2)
	if len(nums) % 2 == 0 {
		// 偶数
		l := len(nums) / 2
		f := l - 1
		return (float64(nums[f]) + float64(nums[l])) / 2
	} else {
		// 奇数
		return float64(nums[(len(nums) - 1) / 2])
	}
}

func main() {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}
	ret := findMedianSortedArrays(nums1, nums2)
	fmt.Println(ret)
}
