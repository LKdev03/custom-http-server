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

	fmt.Printf("\n%v\n", string(temp))

	//return bytes that are not part of this current http request (are part of the next http request or nonimportant)
	if counter < len(temp) {
		return temp[counter+1:]
	} else {
		return []byte{}
	}

}

func handleConnection(conn net.Conn, goroutineId *atomic.Int32) {

	bytes := make([]byte, 2048)

	goroutineId.Add(1)

	//read bytes appear in bytes slice. n is the integer amount of read bytes for this call
	n, err := conn.Read(bytes)
	if err != nil {
		fmt.Printf("error reading bytes from the TCP connection: %v for %v goroutine\n", err, goroutineId)
		return
	}

	fmt.Printf("read %v amount of bytes from TCP connection\n", n)

	fmt.Printf("%v\n", bytes[0:n])
	// bytes -> [255, 255, 255, 255, 123, 123, 321]
	// 			  --  --   --    --

	// add a channel later for giving the return variable to another goroutine of this function
	parseBytes(bytes)

	// n, err := c.Read(b)

}
