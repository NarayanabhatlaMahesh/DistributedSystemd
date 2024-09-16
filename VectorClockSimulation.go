package main

import (
	"fmt"
	"sync"
	"time"
)

// Process represents a distributed process with a vector clock
type Process struct {
	name      string
	vector    []int        // Vector clock
	index     int          // Index of the process in the system
	delay     time.Duration
	sendChans []chan []int // Channels to send messages
	recvChan  chan []int   // Channel to receive messages
}

func (p *Process) InternalEvent(eventName string) {
	// Increment the process's own clock
	p.vector[p.index]++
	fmt.Printf("Process %s: %s occurred at vector time %v\n", p.name, eventName, p.vector)
}

func (p *Process) SendMessage(target int) {
	// Increment the process's own clock before sending a message
	p.vector[p.index]++
	// Send the current vector clock to the target process
	p.sendChans[target] <- append([]int{}, p.vector...) // Send a copy to avoid race conditions
	fmt.Printf("Process %s sent vector %v at vector time %v\n", p.name, p.vector, p.vector)
}

func (p *Process) ReceiveMessage() {
	select {
	case receivedVector := <-p.recvChan:
		// Increment the process's own clock
		p.vector[p.index]++
		// Merge the received vector clock into the process's vector clock
		for i := range p.vector {
			if receivedVector[i] > p.vector[i] {
				p.vector[i] = receivedVector[i]   // we are updating the event value recieved in current vector 
			}
		}
		fmt.Printf("Process %s received vector %v and updated to %v\n", p.name, receivedVector, p.vector)
	default:
	}
}

func (p *Process) Run(wg *sync.WaitGroup, target int) {
	defer wg.Done()

	eventCount := 1
	for eventCount <= 4 {
		if eventCount%2 == 0 {
			p.SendMessage(target)
		} else {
			p.InternalEvent(fmt.Sprintf("e%d", eventCount))
		}
		p.ReceiveMessage()
		eventCount++
		time.Sleep(p.delay)
	}
}

func main() {
	// Channels for message passing
	ch1 := make(chan []int, 1)
	ch2 := make(chan []int, 1)
	ch3 := make(chan []int, 1)

	// Initialize processes with vector clocks
	p1 := &Process{name: "P1", vector: []int{0, 0, 0}, index: 0, delay: 5 * time.Second, sendChans: []chan []int{ch2, ch3}, recvChan: ch1} // recieves from ch1, sends to ch2 & ch3
	p2 := &Process{name: "P2", vector: []int{0, 0, 0}, index: 1, delay: 3 * time.Second, sendChans: []chan []int{ch1, ch3}, recvChan: ch2} // recieves from ch2, sends to ch3 & ch1
	p3 := &Process{name: "P3", vector: []int{0, 0, 0}, index: 2, delay: 7 * time.Second, sendChans: []chan []int{ch1, ch2}, recvChan: ch3} // recieves from ch3, sends to ch2 & ch1

	var wg sync.WaitGroup
	wg.Add(3)

	go p1.Run(&wg, 0) // p1 sends to p2 (index 0)
	go p2.Run(&wg, 1) // p2 sends to p3 (index 1)
	go p3.Run(&wg, 0) // p3 sends to p1 (index 0)

	wg.Wait()
}
