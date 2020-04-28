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

	go func() {
		// read data from a channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()

	go func() {
		i := 42
		// write data to a channel
		ch <- i
		// i value is alredy written to the channel
		// new assigned value will not affect the ch value
		i = 23
		wg.Done()
	}()
	wg.Wait()
}
