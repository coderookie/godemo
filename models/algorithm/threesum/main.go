package main

import (
	"fmt"
	"strconv"
)

func quickSort(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	ret := make([]int, 0)
	f := nums[0]
	l := make([]int, 0)
	s := make([]int, 0)
	for i := 1; i < len(nums); i++ {
		if nums[i] >= f {
			l = append(l, nums[i])
		} else {
			s = append(s, nums[i])
		}
	}

	if len(s) > 0 {
		ret = append(ret, quickSort(s)...)
	}
	ret = append(ret, f)
	if len(l) > 0 {
		ret = append(ret, quickSort(l)...)
	}

	return ret
}

func threeSum(nums []int) [][]int {
	// 从小到大排序
	arr := quickSort(nums)
	ret := make([][]int, 0)

	for i := 0; i < len(arr); i++ {
		// 当前值已经大于0, 再循环下去只会更大, 直接跳出
		if arr[i] > 0 {
			break
		} else if i > 0 && arr[i] == arr[i - 1] {
			// 当前数跟前一个一样, 直接continue
			continue
		}
		for j := i + 1; j < len(arr); j++ {
			if arr[i] + arr[j] > 0 {
				break
			}
			for x := len(arr) - 1; x > j; x-- {
				if arr[i] + arr[j] + arr[x] == 0 {
					ret = append(ret, []int{arr[i], arr[j], arr[x]})
					break
				} else if arr[i] + arr[j] + arr[x] < 0 {
					// 如果v1 + v2 + v3 < 0, 第三次循环便不需要再往前循环了, 往前循环只会更小
					break
				}
			}
		}
	}

	if len(ret) == 0 {
		return ret
	}

	// 去重
	res := make([][]int, 0)
	m := make(map[string]bool)
	for _, item := range ret {
		s := ""
		for _, v := range item {
			t := strconv.Itoa(v)
			s += t
		}
		if _, ok := m[s]; ok {
			// do nothing
		} else {
			res = append(res, item)
			m[s] = true
		}
	}

	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	ret := threeSum(nums)
	fmt.Println(ret)
}
