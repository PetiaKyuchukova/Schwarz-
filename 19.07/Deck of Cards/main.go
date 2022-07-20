package main

import "fmt"

func main() {
	deck := []int{1, 2, 3, 4, 4, 3, 2, 1}

	var exist bool
	var result bool
	sizeX := 0
	indexes := make([]int, 0, len(deck))

	for i := 0; i < len(deck); i++ {
		group := make([]int, 0, len(deck))
		exist = false

		for j := i; j < len(deck); j++ {
			for k := 0; k < len(indexes); k++ {
				if indexes[k] == j {
					exist = true
				}
			}
			if exist == true {
				continue
			}

			if deck[i] == deck[j] {
				group = append(group, deck[j])
				indexes = append(indexes, j)

			}

		}

		if i == 0 {
			sizeX = len(group)
		}

		if len(group) > 0 && len(group) != sizeX {
			result = false
		} else {
			result = true
		}

		if len(group) > 0 {
			fmt.Print(group)
		}

	}
	fmt.Print(result)
}
