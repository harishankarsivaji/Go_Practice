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
	fmt.Println("Server started ...")

	svr, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		panic(err)
	}
	defer svr.Close()

	fmt.Println("Listening on " + HOST + ":" + PORT)
	fmt.Println("Waiting client...")

	for {
		conn, err := svr.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("client connected")
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {

	buff := make([]byte, 1024)
	msg, err := conn.Read(buff)
	if err != nil {
		panic(err)
	}

	fmt.Println("Received: ", string(buff[:msg]))
	conn.Close()
}
