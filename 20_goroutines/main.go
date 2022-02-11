package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var lock = sync.RWMutex{} // create a lock, so only 1 thing can manipulate data at a time
var counter = 0

func main() {
	name := "Jack"

	wg.Add(1)         // add 1 to waitgroup
	go sayHello(name) // start a goroutine (thread)
	wg.Wait()         // wait until wg = 0 (complete all the routines)

	for i := 0; i < 10; i++ {
		wg.Add(2) // add 2 routines to waitgroup

		lock.RLock() // when you rlock a chunck of code
		// it means that any variables used before the lock is unlocked
		// is not allowed to be changed by any other code
		// we lock it here (and not inside the functions), since
		// this will allow the main function (the one running the routines)
		// to set the locks, so there's cannot be 2 printCounter() running before
		// the incrementCounter has ran
		go printCounter()

		lock.Lock() // here we use normal lock, to keep other pieces of code
		// from mutating/reading the variable while we use it
		go incrementCounter()
	}

	wg.Wait()
}

func sayHello(name string) {
	fmt.Println("Hello", name)

	wg.Done() // remove 1 from waitgroup
}

func printCounter() {
	fmt.Println("Counter is", counter)
	lock.RUnlock() // we unlock INSIDE the routine, not outside
	wg.Done()
}

func incrementCounter() {
	counter++
	lock.Unlock() // We unlock INSIDE the routine/function NOT outside
	wg.Done()
}
