package main

import (
	"fmt"
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	ret := ""
	for i := 0; i < len(s); i++ {
		f := string(s[i])
		for j := i + 1; j < len(s); j++ {
			if strings.Contains(f, string(s[j])) {
				break
			} else {
				f += string(s[j])
			}
		}
		if len(f) >= len(ret) {
			ret = f
		}
	}

	return len(ret)
}

func main() {
	s := "bbbb"
	ret := lengthOfLongestSubstring(s)
	fmt.Println(ret)
}