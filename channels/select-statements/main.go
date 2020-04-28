// Go Logger example

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	time     time.Time
	severity string
	message  string
}

var wg = sync.WaitGroup{}
var logCh = make(chan logEntry, 50)

func main() {
	go logger()
	wg.Add(3)
	//write log to the channel
	logCh <- logEntry{time.Now(), logInfo, "Application is starting \n"}

	logCh <- logEntry{time.Now(), logWarning, "Warning message \n"}

	logCh <- logEntry{time.Now(), logError, "Application is shuting down"}
	wg.Wait()
}

func logger() {
	for entry := range logCh {
		fmt.Printf("%v - [%v] - %v ", entry.time.Format("2006-01-02T15:04:05"), entry.severity, entry.message)
		wg.Done()
	}
}
