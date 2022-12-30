package main

import (
	"fmt"
	"math"
)

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	area := 0
	for i := 0; i < len(height); i++ {
		t := 0
		for j := i + 1; j < len(height); j++ {
			t = (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
			if t > area {
				area = t
			}
		}
	}

	return area
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	// height := []int{1, 1}
	ret := maxArea(height)
	fmt.Println(ret)
}
