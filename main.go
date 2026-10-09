package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("hello world")

	Listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error creating listener: %v\n", err)
		return
	}

	for {
		connection, err := Listener.Accept()
		if err != nil {
			fmt.Println("Error accepting TCP connection from OS TCP queue: %v\n", err)
			continue
		}

		// remove this and test speed without concurrency later. Just for fun
		go handleConnection(connection)

	}

}
