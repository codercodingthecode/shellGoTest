package main

import (
	"sync"
	"time"
)

// TickTock time tick
func TickTock(breaker chan bool, wg *sync.WaitGroup, tk chan int) {
Loop:
	for {
		select {
		case <-breaker:
			wg.Done()
			break Loop

		default:
			if time.Now().Second() == 0 && time.Now().Minute() == 0 {
				tk <- 2 // bong
			} else if time.Now().Second() == 0 {
				tk <- 1 // tock
			} else {
				tk <- 0 // tick
			}
		}
	}

}
