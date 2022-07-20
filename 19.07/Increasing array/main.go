package main

import (
	"fmt"
)

func main() {

	nums := [...]int{1, 2, 10, 5, 7}
	canBeIncreased := true

	for i := 0; i < len(nums); i++ {
		checkForIncreasing := [5]int{}
		result := make([]int, 0, 4)

		for j := 0; j < len(checkForIncreasing); j++ {
			if j == i {
				continue
			} else {
				checkForIncreasing[j] = nums[j]
			}
		}

		for k := 0; k < len(checkForIncreasing); k++ {
			if checkForIncreasing[k] != 0 {
				result = append(result, checkForIncreasing[k])
			}
		}
		fmt.Println(result)

		for l := 1; l < len(result); l++ {
			if result[l-1] < result[l] {
				canBeIncreased = true
			} else {
				canBeIncreased = false
				break
			}
		}
		fmt.Println(canBeIncreased)

	}

}
