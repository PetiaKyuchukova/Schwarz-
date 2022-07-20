package main

import (
	"fmt"
	"math"
)

func main() {
	nums := [...]int{8, 9}
	max := math.MinInt
	maximums := make([]int, 0, 3)

	for i := 0; i < 3; i++ {
		max = math.MinInt

		for j := 0; j < len(nums); j++ {
			if i == 0 {
				if nums[j] > max {
					max = nums[j]
				}

			} else if nums[j] > max && nums[j] < maximums[len(maximums)-1] {
				max = nums[j]
			}
		}
		if max != math.MinInt {
			maximums = append(maximums, max)
		}
	}
	max = math.MinInt
	fmt.Println(maximums)

	if len(maximums) >= 3 {
		for i := 0; i < len(maximums); i++ {
			if i == 0 {
				fmt.Println("The first distinct maximum is", maximums[i])
			} else if i == 1 {
				fmt.Println("The second distinct maximum is", maximums[i])
			} else if i == 2 {
				fmt.Println("The third distinct maximum is", maximums[i])
			}
		}
	} else {

		for i := 0; i < len(maximums); i++ {
			if maximums[i] > max {
				max = maximums[i]
			}
		}
		fmt.Println("The third distinct maximum does not exist, so the maximum is", max)
	}

}
