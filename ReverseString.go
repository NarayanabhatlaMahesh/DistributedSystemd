package main

import (
	"fmt"
)

func reverseString(input string) string {
	runes := []rune(input)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

func main() {
	var userInput string
	fmt.Println("Enter a string to reverse:")
	fmt.Scanln(&userInput)
	reversedString := reverseString(userInput)
	fmt.Println("Reversed string:", reversedString)
}
