package main

import (
	"fmt"
)

func permute1(d [][][]int, l int) [][][]int {
	for _, v := range d {
		if len(v[0]) == l {
			return d
		}
	}

	ret := make([][][]int, 0)

	// d = [ [[1], [2, 3]], [[2], [1, 3]], [[3], [1, 2]] ]
	for _, v := range d {
		for i := 0; i < len(v[1]); i++ {
			// f = [1]
			f := make([]int, len(v[0]))
			copy(f, v[0])

			// f = [1, 2]
			f = append(f, v[1][i])

			// le = [3]
			le := make([]int, 0)
			if len(v[1][:i]) > 0 {
				le = append(le, v[1][:i]...)
			}
			if len(v[1][i + 1:]) > 0 {
				le = append(le, v[1][i + 1:]...)
			}

			// e = [[1, 2], [3]]
			e := make([][]int, 0)
			e = append(e, f)
			e = append(e, le)

			// ret = [ [[1, 2], [3]], [[1, 3], [2]], [[2, 1], [3]], [[2, 3], [1]], ... ]
			ret = append(ret, e)
		}
	}

	return permute1(ret, l)
}

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}

	arr := make([][][]int, 0)
	for i := 0; i < len(nums); i++ {
		e := make([][]int, 0)

		f := make([]int, 0)
		f = append(f, nums[i])

		l := make([]int, 0)
		if len(nums[:i]) > 0 {
			l = append(l, nums[:i]...)
		}
		if len(nums[i + 1:]) > 0 {
			l = append(l, nums[i + 1:]...)
		}

		e = append(e, f)
		e = append(e, l)

		arr = append(arr, e)
	}

	ret := make([][]int, 0)
	res := permute1(arr, len(nums))
	for _, v := range res {
		ret = append(ret, v[0])
	}

	return ret
}

func main() {
	nums := []int{1, 2, 3}
	ret := permute(nums)
	fmt.Println(ret)
}
