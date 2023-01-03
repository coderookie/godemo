package main

import (
	"fmt"
)

func isValid(s string) bool {
	m := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	ret := make([]string, 0)
	for _, b := range s {
		b1 := string(b)
		t, _ := m[b1]
		if len(ret) > 0 && ret[len(ret) - 1] == t {
			ret = ret[0:len(ret) - 1]
		} else {
			ret = append(ret, b1)
		}
	}

	if len(ret) == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	// s := "()[]{}"
	s := "{]"
	ret := isValid(s)
	fmt.Println(ret)
}
