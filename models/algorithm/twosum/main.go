package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	m := make(map[int]int)
	for k, v := range nums {
		if k1, ok := m[v]; ok {
			return []int{k1, k}
		} else {
			m[target - v] = k
		}
	}

	return []int{}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	ret := twoSum(nums, target)
	fmt.Println(ret)
}
