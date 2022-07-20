package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "c9t and. do-g"
	countOfValidWords := 0
	countOfHyphen := 0
	countOfPunctuation := 0
	validWord := false

	words := strings.Split(sentence, " ")

	for i := 0; i < len(words); i++ {
		letters := []rune(words[i])
		validWord = true
		countOfHyphen = 0
		countOfPunctuation = 0

		for j := 0; j < len(letters); j++ {

			if !(97 <= letters[j] && letters[j] <= 122) {
				if letters[j] == 45 {
					countOfHyphen++

					if !(97 <= letters[j-1] && letters[j-1] <= 122) ||
						!(97 <= letters[j+1] && letters[j+1] <= 122) || countOfHyphen > 1 {
						validWord = false

						break

					}
				} else if letters[j] == 46 || letters[j] == 63 || letters[j] == 44 || letters[j] == 33 {

					countOfPunctuation++

					if !(letters[j] == letters[len(letters)-1]) || countOfPunctuation > 1 {
						validWord = false
						break
					}

				} else {
					validWord = false
					break
				}
			}

		}

		if validWord == true {
			countOfValidWords++
		}

	}
	fmt.Print(countOfValidWords)

}
