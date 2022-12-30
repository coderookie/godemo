package main

import (
	"fmt"
)

func chkPalindrome(s string) bool {
	ret := true
	for i := 0; i < len(s); i++ {
		if i > len(s) / 2 {
			break
		}
		if string(s[i]) != string(s[len(s) - i - 1]) {
			return false
		}
	}

	return ret
}

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	// 如果全是一样的字符, 直接返回字符
	same := true
	f := string(s[0])
	for i := 0; i < len(s); i++ {
		if string(s[i]) != f {
			same = false
			break
		}
	}
	if same {
		return s
	}

	ret := ""
	for i := 0; i < len(s); i++ {
		f := string(s[i])
		if len(ret) == 0 {
			ret = f
		}
		for j := i + 1; j < len(s); j++ {
			f = f + string(s[j])
			if chkPalindrome(f) && len(ret) < len(f) {
				ret = f
			}
		}
	}

	return ret
}

func main() {
	s := "ac"
	ret := longestPalindrome(s)
	fmt.Println(ret)
}