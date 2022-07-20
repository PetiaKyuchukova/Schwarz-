package main

import (
	"fmt"
)

func main() {
	s := "abcd"
	goal := "dbca"
	indexes := make([]int, 0)
	var result bool

	if len(s) == len(goal) {
		for i := 0; i < len(s); i++ {
			if s[i] != goal[i] {
				indexes = append(indexes, i)
			}
		}
		if len(indexes) == 2 {
			result = true
			fmt.Println(result)
			fmt.Println("You can swap these indexes: ", indexes)
		} else {
			result = false
			fmt.Print(result)

		}
	} else {
		result = false
		fmt.Print("Not Buddy Strings")
	}

}
