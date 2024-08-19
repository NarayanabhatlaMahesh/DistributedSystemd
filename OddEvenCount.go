package main

import (
	"fmt"
)

func countOddEven(numbers []int) (int, int) {
	oddCount := 0
	evenCount := 0
	for _, number := range numbers {
		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	return oddCount, evenCount
}

func main() {
	var arraySize int
	fmt.Println("Enter the number of elements in the array:")
	fmt.Scan(&arraySize)

	numbers := make([]int, arraySize)
	fmt.Println("Enter the elements of the array:")
	for i := 0; i < arraySize; i++ {
		fmt.Scan(&numbers[i])
	}

	oddCount, evenCount := countOddEven(numbers)
	fmt.Printf("Odd numbers: %d, Even numbers: %d\n", oddCount, evenCount)
}
