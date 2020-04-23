package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var m = sync.RWMutex{}
var counter = 0

func main() {
	// runtime.GOMAXPROCS(1)
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello Go #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}

// func main() {
// 	var msg = "Hello"
// 	wg.Add(1)
// 	go func(msg string) {
// 		fmt.Println(msg)
// 		wg.Done()
// 	}(msg)
// 	msg = "Goodbye"
// 	wg.Wait()
// }
