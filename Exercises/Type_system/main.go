package main

import (
	"fmt"
	time2 "time"
)

//package scope
//Live forever
var myAge, yourAge = 50, 2

func main() {

	fmt.Println(
		-200, -100, 0, 50, 100)

	fmt.Println(
		-50.5, -20.5, -0., 1.01, 0x32, 0x64) // last 2 are hexadecimal representations

	fmt.Println(
		true, false)

	fmt.Println(
		"", "Hello There !!!")

	// parallel mulitple declaration

	var temperature float64
	time := time2.Now()
	var brand string
	_ = brand // Blank identifier

	fmt.Println(myAge, yourAge)
	fmt.Println(temperature)
	fmt.Println(time)
}
