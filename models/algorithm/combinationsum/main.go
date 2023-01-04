package main

import (
	"fmt"
	"sort"
	"strconv"
)

func isIn(t int, candidates []int) bool {
	for _, v := range candidates {
		if t == v {
			return true
		}
	}

	return false
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	// 先排序
	sort.Ints(candidates)
	if candidates[0] > target {
		return [][]int{}
	}

	ret := make([][]int, 0)

	for _, v := range candidates {
		// 如果当前值等于target, 则把自己放入返回值
		if v == target {
			ret = append(ret, []int{v})
		}

		// 剩余值 = target - 当前值, 剩余值直接在数组中, 则放入返回值
		t := make([]int, 0)
		t = append(t, v)
		less := target - v
		if isIn(less, candidates) {
			ret = append(ret, []int{v, less})
		}

		// 剩余值不在数组中, 则递归
		r := combinationSum(candidates, less)
		if len(r) > 0 {
			for _, va := range r {
				tm := make([]int, 0)
				tm = append(tm, v)
				tm = append(tm, va...)
				ret = append(ret, tm)
			}
		}
	}

	// 一维数组排序
	if len(ret) == 0 {
		return [][]int{}
	}

	res := make([][]int, 0)
	m := make(map[string]bool)
	for _, v := range ret {
		sort.Ints(v)
		str := ""
		for _, n := range v {
			s := strconv.Itoa(n)
			str += s
		}
		if _, ok := m[str]; !ok {
			res = append(res, v)
			m[str] = true
		}
	}

	return res
}

func main() {
	// candidates := []int{2, 3, 6, 7}
	// target := 7
	candidates := []int{8, 7, 4, 3}
	target := 11
	ret := combinationSum(candidates, target)
	fmt.Println(ret)
}
