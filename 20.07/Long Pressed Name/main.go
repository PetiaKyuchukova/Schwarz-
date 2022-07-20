package main

import "fmt"

func main() {
	name := "saeed"
	typed := "ssaaedd"
	checkedLetter := make([]string, 0)
	checker := true
	nameCounter := 0
	typedCounter := 0

	for i := 0; i < len(typed); i++ {

		if len(checkedLetter) > 0 && checkedLetter[len(checkedLetter)-1] == string(typed[i]) {
			continue
		}

		checker = false
		nameCounter = 0
		typedCounter = 0
		for j := 0; j < len(name); j++ {
			if typed[i] == name[j] {
				checker = true
				nameCounter++
			}
		}

		if checker == false {
			break
		}
		for l := 0; l < len(typed); l++ {
			if typed[i] == typed[l] {
				typedCounter++
			}
		}

		if typedCounter > nameCounter {
			checkedLetter = append(checkedLetter, string(typed[i]))
			fmt.Println(string(typed[i]), "is long press")
		} else if typedCounter < nameCounter {
			fmt.Println(string(typed[i]), "must have been pressed", nameCounter, "times")

		}

	}
	fmt.Println("Output:", checker)
}
