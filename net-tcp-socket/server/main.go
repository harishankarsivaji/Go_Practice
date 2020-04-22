/* Server socket application to listen incominng client msg
* Listen() - creates the server with type of connection and on specific port
* Accept() - accepts connection from the client
* https://golang.org/pkg/net/
 */

package main

import (
	"fmt"
	"net"
)

const (
	PORT = "9092"
	TYPE = "tcp"
)

func main() {
	fmt.Println("Server started")

	svr, err := net.Listen(TYPE, ":"+PORT)
	if err != nil {
		panic(err)
	}
	defer svr.Close()

	fmt.Println("Listening on localhost" + ":" + PORT)
	fmt.Println("Waiting client...")

	for {
		conn, err := svr.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Client connected")
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
