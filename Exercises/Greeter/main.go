package main

import (
	"fmt"
	"os"
)

func main() {
	//variable()
	name := os.Args[1]

	fmt.Println("Welcome !!!", name)

	bot := "Harish"

	fmt.Println("My name is", bot, "How are you today?")

	fmt.Println("I am your chat bot")
}
