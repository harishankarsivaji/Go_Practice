// Buffered Cahnnel 
// Closing channel 
// For .. range loop with channel 

ackage main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	// define a channel with buffer (internal data store of channel)
	ch := make(chan int, 50)
	wg.Add(2)

	//receive-only channel
	go func(ch <-chan int) {
		for i := range ch {
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	//send-only channel
	go func(ch chan<- int) {
		// write data to a channel
		ch <- 42
		ch <- 27
		// when using for..range with channel, close the channel to avoid deadlock
		close(ch)
		wg.Done()
	}(ch)
	wg.Wait()
}
