package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// channels are what we use to send data between go routines

	channel := make(chan int)
	// EXAMPLE 1
	wg.Add(2)

	// receive only channel (only allowed to receive/read data)
	// we force it to be receive only by adding the
	// "ch <-chan int" parameter (optional)
	go func(ch <-chan int) {
		i := <-ch // receive data from channel
		fmt.Println(i)
		wg.Done()
	}(channel)

	// send only channel (only allowed to send/change data)
	// we force it to be send only by adding the
	// "ch chan<- int" parameter (optional)
	go func(ch chan<- int) {
		ch <- 42 // send data into a channel
		wg.Done()
	}(channel)

	wg.Wait()

	// EXAMPLE 2
	wg.Add(2)

	// receive only channel (only allowed to receive/read data)
	// we force it to be receive only by adding the
	// "ch <-chan int" parameter (optional)
	go func(ch <-chan int) {
		for i := range ch {
			// to receive all the channel
			fmt.Println(i)
		}
		wg.Done()
	}(channel)

	// send only channel (only allowed to send/change data)
	// we force it to be send only by adding the
	// "ch chan<- int" parameter (optional)
	go func(ch chan<- int) {
		ch <- 42 // send data into a channel
		ch <- 3
		ch <- 20
		close(ch) //close the channel to tell range ch that all data has been sent
		// NOTE: When you close a channel, you can not use it again!
		// so if I close the channel in example 1, then this example
		// won't work
		wg.Done()
	}(channel)

	wg.Wait()
}
