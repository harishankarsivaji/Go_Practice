package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// define a channel
	ch := make(chan int)
	wg.Add(2)

	//receive-only channel
	go func(ch <-chan int) {
		// read data from a channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	//send-only channel
	go func(ch chan<- int) {
		i := 42
		// write data to a channel
		ch <- i
		wg.Done()
	}(ch)
	wg.Wait()
}
