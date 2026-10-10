package main

import (
	"fmt"
	"net"
	"sync/atomic"
)

// parse bytes from given buffer, find delimiters and create http request
func parseBytes(buffer []byte) []byte {

	temp := buffer

	//counter for counting how many bytes the current http call was
	counter := 0

	//return bytes that are not part of this current http request

	if counter < len(temp) {
		return temp[counter+1:]
	} else {
		return []byte{}
	}

}

func handleConnection(conn net.Conn, goroutineId *atomic.Int32) {

	bytes := []byte{}

	goroutineId.Add(1)

	//read bytes appear in bytes slice. n is the integer amount of read bytes for this call
	n, err := conn.Read(bytes)
	if err != nil {
		fmt.Printf("error reading bytes from the TCP connection: %v for %v goroutine", err, goroutineId)
		return
	}

	// bytes -> [255, 255, 255, 255, 123, 123, 321]
	// 			  --  --   --    --

	bytes = parseBytes(bytes)

	// n, err := c.Read(b)

}
