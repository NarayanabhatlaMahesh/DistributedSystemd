package main

import (
	"fmt"
)

func checkRollNumber(rollNumber int) {
	lastThreeDigits := rollNumber % 1000
	if lastThreeDigits%2 == 0 {
		fmt.Println("Your last name is: Mahesh") // Replace with your last name
	} else {
		fmt.Println("Your first name is: Mahesh") // Replace with your first name
	}
}

func main() {
	var rollNumber int
	fmt.Println("Enter your roll number:")
	fmt.Scan(&rollNumber)
	checkRollNumber(rollNumber)
}
