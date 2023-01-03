package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
	}
	ret := make([]string, 0)
	for _, b := range s {
		t := string(b)
		e, ok := m[t]
		if len(ret) > 0 && ok && ret[len(ret) - 1] == e {
			ret = ret[:len(ret) - 1]
		} else {
			ret = append(ret, t)
		}
	}

	if len(ret) == 0 {
		return true
	}

	return false
}

func pailie(d [][][]string, l int) [][]string {
	over := false
	ret := make([][]string, 0)
	for _, v := range d {
		if len(v[0]) == l {
			over = true
			ret = append(ret, v[0])
		}
	}
	if over {
		return ret
	}

	// d = [ [["("], [")", "(", ")", "(", ")"]] ]
	s := make([][][]string, 0)
	m := make(map[string]bool)
	for _, v := range d {
		for i := 0; i < len(v[1]); i++ {
			t := make([][]string, 0)

			// fi = ["("], le = []
			fi := make([]string, len(v[0]))
			le := make([]string, 0)
			copy(fi, v[0])
			fi = append(fi, v[1][i])
			if len(v[1][:i]) > 0 {
				le = append(le, v[1][:i]...)
			}
			if len(v[1][i + 1:]) > 0 {
				le = append(le, v[1][i + 1:]...)
			}

			t = append(t, fi)
			t = append(t, le)

			d := make([]string, 0)
			d = append(d, fi...)
			d = append(d, le...)
			_, ok := m[strings.Join(d, "")]
			if !ok {
				m[strings.Join(d, "")] = true
				s = append(s, t)
			}
		}
	}

	return pailie(s, l)
}

func generateParenthesis(n int) []string {
	// s = ["(", ")", "(", ")", "(", ")"]
	s := make([]string, 0)
	for i := 1; i <= n; i++ {
		s = append(s, "(")
		s = append(s, ")")
	}

	// arr = [ [["("], [")", "(", ")", "(", ")"]], [[[")"], ["(", "(", ")", "(", ")"]], [["("], [")", "(", ")", "(", ")"]]... ]
	arr := make([][][]string, 0)
	dup := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		// f = ["("]
		f := make([]string, 0)
		f = append(f, s[i])

		// l = [")", "(", ")", "(", ")"]
		l := make([]string, 0)
		if len(s[0:i]) > 0 {
			l = append(l, s[0:i]...)
		}
		if len(s[i + 1:]) > 0 {
			l = append(l, s[i + 1:]...)
		}

		// e = [["("], [")", "(", ")", "(", ")"]]
		e := make([][]string, 0)
		e = append(e, f)
		e = append(e, l)

		tm := make([]string, 0)
		tm = append(tm, f...)
		tm = append(tm, l...)
		for _, va := range tm {
			_, ok := dup[va]
			if !ok {
				arr = append(arr, e)
				dup[va] = true
			}
		}
	}

	ret := make([]string, 0)
	res := pailie(arr, 2 * n)
	for _, v := range res {
		if isValid(strings.Join(v, "")) {
			ret = append(ret, strings.Join(v, ""))
		}
	}

	// 去重
	m := make(map[string]bool)
	re := make([]string, 0)
	for _, v := range ret {
		_, ok := m[v]
		if !ok {
			m[v] = true
			re = append(re, v)
		}
	}

	return re
}

func main() {
	n := 8
	ret := generateParenthesis(n) // ["((()))","(()())","(())()","()(())","()()()"]
	fmt.Println(ret)
}
