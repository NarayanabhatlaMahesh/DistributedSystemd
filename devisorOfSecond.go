package main

import (
	"fmt"
)

func main() {
	var firstNumber, secondNumber int
	fmt.Println("Enter two numbers:")
	fmt.Scan(&firstNumber, &secondNumber)

	if secondNumber%firstNumber == 0 {
		fmt.Printf("%d is a divisor of %d\n", firstNumber, secondNumber)
	} else {
		fmt.Printf("%d is not a divisor of %d\n", firstNumber, secondNumber)
	}
}
