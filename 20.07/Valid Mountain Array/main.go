package main

import "fmt"

func main() {
	arr := [...]int{2, 1}
	leftSubArray := [len(arr) / 2]int{}
	rightSubArray := [len(arr) / 2]int{}
	var peak int

	validMountainArray := true
	if len(arr) < 3 {
		validMountainArray = false
	} else if len(arr)%2 == 0 {
		for i := 0; i < len(arr); i++ {
			if i <= (len(arr)/2)-1 {
				leftSubArray[i] = arr[i]
			} else {
				rightSubArray[i-len(rightSubArray)] = arr[i]
			}
		}
	} else {
		for i := 0; i < len(arr); i++ {
			if i <= (len(arr)/2)-1 {
				leftSubArray[i] = arr[i]
			} else if i == len(arr)/2 {
				peak = arr[i]
			} else {
				rightSubArray[i-len(rightSubArray)-1] = arr[i]
			}
		}
	}

	for i := 1; i < len(leftSubArray); i++ {
		if !(leftSubArray[i-1] < leftSubArray[i]) {
			validMountainArray = false
			fmt.Println("Left side is not strickly increasing!")
			break
		}
		if !(rightSubArray[i-1] > rightSubArray[i]) {
			validMountainArray = false
			fmt.Println("Right side is not strickly decreasing!")
			break
		}
	}

	if peak > 0 && (peak <= leftSubArray[len(leftSubArray)-1] || peak <= rightSubArray[0]) {
		validMountainArray = false
		fmt.Println("The peak is not the only ant the highest point on the mountain!")
	}

	if validMountainArray == true {
		fmt.Println("Output:", validMountainArray)
		fmt.Println("The aray is a valid mountain array")
	} else {
		fmt.Println("Output:", validMountainArray)
		fmt.Println("The aray is not a valid mountain array")
	}

	fmt.Println("Left side of the mountain is ", leftSubArray)
	fmt.Println("The peak of the mountain is ", peak)
	fmt.Println("Right side of the mountain is ", rightSubArray)

}
