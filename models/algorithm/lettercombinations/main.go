package main

import (
	"fmt"
)

func combinations(n1 []string, n2 []string) []string {
	if len(n1) == 0 {
		return n2
	} else if len(n2) == 0 {
		return n1
	}

	ret := make([]string, 0)
	for _, v1 := range n1 {
		for _, v2 := range n2 {
			ret = append(ret, v1 + v2)
		}
	}

	return ret
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}

	// 生成二维数组
	a := make([][]string, 0)
	for _, b := range digits {
		if v, ok := m[string(b)]; ok {
			a = append(a, v)
		}
	}

	// 返回值
	ret := make([]string, 0)
	for _, v := range a {
		ret = combinations(ret, v)
	}

	return ret
}

func main() {
	digits := "23"
	ret := letterCombinations(digits)
	fmt.Println(ret)
}
