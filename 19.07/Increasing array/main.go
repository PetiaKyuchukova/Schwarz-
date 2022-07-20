package main

import (
	"fmt"
)

func main() {

	nums := [...]int{1, 1, 1}
	canBeIncreased := true

	for i := 0; i < len(nums); i++ {
		checkForIncreasing := [len(nums)]int{}
		result := make([]int, 0, len(nums))

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
