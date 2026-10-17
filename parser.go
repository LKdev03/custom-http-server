/*
curl -X POST https://example.com/api/users \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  -d '{
    "name": "John",
    "email": "john@example.com"
  }'



  ->

POST /api/users HTTP/1.1\r\n
Host: example.com\r\n
Content-Type: application/json\r\n
Accept: application/json\r\n
Content-Length: 49\r\n
\r\n
{
    "name": "John",
    "email": "john@example.com"
  }

*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {

	requestString := "POST /api/users HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"Content-Type: application/json\r\n" +
		"Accept: application/json\r\n" +
		"Content-Length: 21\r\n" +
		"\r\n" +
		`{"name":"John"}`

	slice := []byte(requestString)

	// TODO:
	// 1. check for Content-Length and Transfer-Encoding
	// 2. if either of them exists, parse the value
	// 3. Find the final delimiter: \r\n\r\

	isContentLength := bytes.Contains(slice, []byte("Content-Length"))
	isTransferEncoding := bytes.Contains(slice, []byte("Transfer-Encoding"))

	if isContentLength {

		//Parse Content-Length
		copy := slice

		//cut the bytes slice to when content-length is found and take only what happens after.
		_, after, _ := bytes.Cut(copy, []byte("Content-Length: "))

		contentLength := []byte{}

		for i := 0; i < len(after); i++ {
			// check if we have iterated over the content-length variable
			// 13 equals '\r' in ascii
			// 10 equals '\n' in ascii

			if after[i] == 13 {
				break
			}

			contentLength = append(contentLength, []byte(after)[i])
		}
		fmt.Println(contentLength)

		contentLengthInt, err := strconv.Atoi(string(contentLength))
		if err != nil {
			//handle error, not necessary for this test script
		}

		fmt.Printf("contentLengthInt value is: %v and type is: %T\n", contentLengthInt, contentLengthInt)

	} else if isTransferEncoding {
		//Parse Transfer Encoding

		//TODO
	}

	//fmt.Println(slice)

}
