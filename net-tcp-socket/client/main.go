/* Client socket application to send meassge to server
* Dial() - to connect to server
* Write() - to send data
 */

package main

import (
	"fmt"
	"net"
)

const (
	// PORT is a constant
	PORT = "9092"
	// TYPE is a constant
	TYPE = "tcp"
)

func main() {
	fmt.Println("Client started ")
	conn, err := net.Dial(TYPE, ":"+PORT)
	if err != nil {
		panic(err)
	}

	fmt.Println("Sending data ...")
	_, err = conn.Write([]byte("Hello from client !!"))
	defer conn.Close()
}
