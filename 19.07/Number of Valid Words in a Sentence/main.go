package main

import (
	"fmt"
	"strings"
)

func main() {

	sentence := "cat and  dog"
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
			if 97 <= letters[j] && letters[j] <= 122 {
				validWord = true
			} else if len(letters) == 0 {
				validWord = false
				break
			} else {
				validWord = false
				if letters[j] == 45 {
					countOfHyphen++
					validWord = true

					if j == 0 {
						validWord = false
						break
					} else if j == len(letters)-1 {
						validWord = false
						break
					} else if !(97 <= letters[j-1] && letters[j-1] <= 122) {
						validWord = false
						break
					} else if !(97 <= letters[j+1] && letters[j+1] <= 122) {
						validWord = false
						break
					} else if countOfHyphen > 1 {
						validWord = false
						break
					}

				} else if letters[j] == 46 || letters[j] == 63 || letters[j] == 44 || letters[j] == 33 {
					validWord = true
					countOfPunctuation++
					if j == 0 {
						validWord = false
						break
					} else if countOfPunctuation > 1 {
						validWord = false
						break
					} else if !(j == len(letters)-1) {
						validWord = false
						break
					}

				} else {
					validWord = false
					break
				}
			}

		}

		if validWord == true && len(words[i]) > 0 {
			countOfValidWords++
		}

	}
	fmt.Print(countOfValidWords)

}
