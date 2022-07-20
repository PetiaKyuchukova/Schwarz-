package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [3]int{3, 2, 2}
	max := math.MinInt

	for i := 0; i < 3; i++ {
		max = math.MinInt

		for j := i; j < len(nums); j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}

		fmt.Print(max)
	}
}
