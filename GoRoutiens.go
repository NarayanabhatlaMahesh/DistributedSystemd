package main

import (
	"fmt"
	"sync"
)

func reverseArray(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for left, right := 0, len(numbers)-1; left < right; left, right = left+1, right-1 {
		numbers[left], numbers[right] = numbers[right], numbers[left]
	}
	fmt.Println("Reversed array:", numbers)
}

func countOddEvenElements(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	oddCount := 0
	evenCount := 0
	for _, number := range numbers {
		if number%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Printf("Odd numbers: %d, Even numbers: %d\n", oddCount, evenCount)
}

func countZeros(numbers []int, wg *sync.WaitGroup) {
	defer wg.Done()
	zeroCount := 0
	for _, number := range numbers {
		if number == 0 {
			zeroCount++
		}
	}
	fmt.Printf("Number of zeros: %d\n", zeroCount)
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

	var wg sync.WaitGroup

	wg.Add(3)
	go reverseArray(numbers, &wg)
	go countOddEvenElements(numbers, &wg)
	go countZeros(numbers, &wg)

	wg.Wait()
}
