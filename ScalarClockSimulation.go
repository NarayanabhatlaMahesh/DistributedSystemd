package main

import (
	"fmt"
	"sync"
	"time"
)

type Process struct {
	name      string
	clock     int
	delay     time.Duration
	sendChans []chan string 
	recvChan  chan string   
}

func (p *Process) InternalEvent(eventName string) {
	p.clock++ 
	fmt.Printf("Process %s: %s occurred at scalar time %d\n", p.name, eventName, p.clock)
}

func (p *Process) SendMessage(target int, message string) {
	p.clock++ 
	p.sendChans[target] <- message
	fmt.Printf("Process %s sent message '%s' at scalar time %d\n", p.name, message, p.clock)
}

func (p *Process) ReceiveMessage() {
	select {
	case message := <-p.recvChan:
		p.clock++ 
		fmt.Printf("Process %s received message '%s' at scalar time %d\n", p.name, message, p.clock)
	default:
	}
}

func (p *Process) Run(wg *sync.WaitGroup, target int) {
	defer wg.Done() 

	eventCount := 1
	for eventCount <= 4 { 
		if eventCount%2 == 0 { 
			p.SendMessage(target, fmt.Sprintf("Message from %s", p.name))
		} else {
			p.InternalEvent(fmt.Sprintf("e%d", eventCount)) 
		}
		p.ReceiveMessage() 
		eventCount++
		time.Sleep(p.delay)
	}
}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	p1 := &Process{name: "P1", clock: 0, delay: 5 * time.Second, sendChans: []chan string{ch2, ch3}, recvChan: ch1} // recieves from ch1, sends to ch2 & ch3
	p2 := &Process{name: "P2", clock: 0, delay: 3 * time.Second, sendChans: []chan string{ch1, ch3}, recvChan: ch2} // recieves from ch2, sends to ch3 & ch1
	p3 := &Process{name: "P3", clock: 0, delay: 7 * time.Second, sendChans: []chan string{ch1, ch2}, recvChan: ch3} // recieves from ch3, sends to ch2 & ch1

	var wg sync.WaitGroup
	wg.Add(3) 
	go p1.Run(&wg, 0) // p1 sends to p2 (index 0)
	go p2.Run(&wg, 1) // p2 sends to p3 (index 1)
	go p3.Run(&wg, 0) // p3 sends to p1 (index 0)

	wg.Wait()
}
