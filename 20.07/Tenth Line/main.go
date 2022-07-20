package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("text.txt")
	n := 0

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		n++
		if n < 0 {
			continue
		}
		if n > 10 {
			break
		}
		fmt.Println(fileScanner.Text())
	}

	readFile.Close()
}
