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
	"errors"
	"fmt"
	"strconv"
)

type Request struct {
	//request line
	Host        string
	Method      string
	Target      string
	HttpVersion string

	//Headers
	Connection       string
	ContentType      string
	Accept           string
	ContentLength    int
	TransferEncoding int

	//Body
	Body string
}

// Currently assumes the request does not come in chunks (not how TCP works).
// TODO: change to handle TCP packets instead.

// -1 error
//
//	1  content-length exists
//	2  Transfer-Encoding exists
//	3  neither exists exists
func parseRequestLengthHeaders(request *Request) (int, error) {

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

		contentLengthInt, err := strconv.Atoi(string(contentLength))
		if err != nil {
			//Cant get content-length value
			return -1, errors.New("Cant parse content-length value\n")
		}

		fmt.Printf("contentLengthInt value is: %v and type is: %T\n", contentLengthInt, contentLengthInt)

		request.ContentLength = contentLengthInt

		return 1, nil

	} else if isTransferEncoding {
		//Parse Transfer Encoding

		//TODO

		return 2, nil
	}

	return 3, nil

}

func main() {

	request := Request{}

	//parseBodyFlag, err := parseRequestLengthHeaders(&request)
	_, err := parseRequestLengthHeaders(&request)
	if err != nil {
		//
		return
	}

}
