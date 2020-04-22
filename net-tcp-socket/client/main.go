package main

import (
	"fmt"
	"net"
)

const (
	HOST = "localhost"
	PORT = "1992"
	TYPE = "tcp"
)

func main() {
	fmt.Println("Client started ...")
	conn, err := net.Dial(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending data")
	_, err = conn.Write([]byte("Hello from client !!"))
	defer conn.Close()
}
