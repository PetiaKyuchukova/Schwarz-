package main

import "fmt"

func main() {
	flowerbed := [...]int{0, 0, 0, 0, 1}
	n := 2
	var possiblePlaces int

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] != 0 && flowerbed[i] != 1 {
			fmt.Println("Invalid element!")
			break
		} else {
			if i == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				possiblePlaces++
			} else if i == len(flowerbed)-1 && flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				possiblePlaces++
			} else if i > 0 && flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				possiblePlaces++
			}
		}
	}
	fmt.Println(possiblePlaces)
	fmt.Println(flowerbed)
	if n == possiblePlaces {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}
