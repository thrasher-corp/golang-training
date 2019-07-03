package main

import "fmt"

var number = 10

func main() {
	if isEven(number) {
		fmt.Printf("Number %d is even", number)
	}

	if isOdd(number) {
		fmt.Printf("Number %d is odd", number)
	}
}

func isEven(i int) bool {
	return i%2 == 0
}

func isOdd(i int) bool {
	return i%2 != 0
}
