package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	start := -1
	end := -1
	for i := 0; i < len(nums); i++ {
		if start == -1 && nums[i] == target {
			start = i
			end = i
		} else if nums[i] == target {
			end = i
		}
	}

	return []int{start, end}
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	ret := searchRange(nums, target)
	fmt.Println(ret)
}