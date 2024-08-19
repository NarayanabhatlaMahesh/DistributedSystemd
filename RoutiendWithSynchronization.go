package main

import (
	"fmt"
	"sync"
)

func getName(nameChannel chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var userName string
	fmt.Println("Enter your name:")
	fmt.Scan(&userName)
	nameChannel <- userName
}

func getRollNumber(rollNumberChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	var rollNumber int
	fmt.Println("Enter your roll number:")
	fmt.Scan(&rollNumber)
	rollNumberChannel <- rollNumber
}

func printData(nameChannel chan string, rollNumberChannel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	userName := <-nameChannel
	rollNumber := <-rollNumberChannel
	fmt.Printf("Name: %s, Roll Number: %d\n", userName, rollNumber)
}

func main() {
	var wg sync.WaitGroup
	nameChannel := make(chan string)
	rollNumberChannel := make(chan int)

	wg.Add(3)
	go getName(nameChannel, &wg)
	go getRollNumber(rollNumberChannel, &wg)
	go printData(nameChannel, rollNumberChannel, &wg)

	wg.Wait()
}
