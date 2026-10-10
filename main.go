package main

import (
	"fmt"
	"net"
	"sync/atomic"
)

func main() {
	fmt.Println("hello world from main function")

	Listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error creating listener: %v\n", err)
		return
	}

	for {
		connection, err := Listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting TCP connection from OS TCP queue: %v\n", err)
			continue
		}

		var goroutineId atomic.Int32

		// remove this and test speed without concurrency later. Just for fun
		go handleConnection(connection, &goroutineId)

	}

}
