package main

import (
	"fmt"
)

// Function to get the name based on roll number
func getNameByRollNumber(rollNumber int, rollNumberMap map[int]string) string {
	if name, exists := rollNumberMap[rollNumber]; exists {
		return name
	}
	return "Roll number not found"
}

// Function to add a new roll number and name to the map
func addNewStudent(rollNumber int, name string, rollNumberMap map[int]string) int {
	rollNumberMap[rollNumber] = name
	return rollNumber + 1 // Increment the roll number for the next student
}

func main() {
	// Initial roll number
	currentRollNumber := 24015

	// Map of roll numbers to names
	rollNumberMap := make(map[int]string)

	// Adding students to the map
	currentRollNumber = addNewStudent(currentRollNumber, "Mahesh", rollNumberMap)
	currentRollNumber = addNewStudent(currentRollNumber, "Anita", rollNumberMap)
	currentRollNumber = addNewStudent(currentRollNumber, "Rahul", rollNumberMap)
	currentRollNumber = addNewStudent(currentRollNumber, "Sneha", rollNumberMap)

	// User inputs roll number to get the associated name
	var rollNumber int
	fmt.Println("Enter your roll number:")
	fmt.Scan(&rollNumber)

	// Get the name for the given roll number
	name := getNameByRollNumber(rollNumber, rollNumberMap)
	fmt.Println("The name corresponding to the roll number is:", name)
}
