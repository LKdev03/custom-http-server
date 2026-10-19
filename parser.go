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
	"reflect"
	"strings"
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
	ContentLength    string // int
	TransferEncoding string // int

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
/* func parseRequestLengthHeaders(request *Request) (int, error) {

	//example.com/user/getId

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

} */

//func parseField(string, )

func parseString(str []byte, request *Request) error {

	copy := str

	//Create a hashmap for turning raw input to struct fields:
	// Content-Length -> ContentLength

	transformations := map[string]string{
		"Content-Length":    "ContentLength",
		"Content-Type":      "ContentType",
		"User-Agent":        "UserAgent",
		"Accept-Language":   "AcceptLanguage",
		"Accept-Encoding":   "AcceptEncoding",
		"Cache-Control":     "CacheControl",
		"If-None-Match":     "IfNoneMatch",
		"Transfer-Encoding": "TransferEncoding",
	}

	before, after, _ := bytes.Cut(copy, []byte(":"))

	fmt.Printf("before: %v and after: %v\n", string(before), string(after))

	fieldName := string(before)

	value, ok := transformations[string(before)]

	if ok {
		fieldName = value
	}

	//check if the field exists in the request body struct
	t := reflect.ValueOf(request).Elem()
	field := t.FieldByName(fieldName)

	if !field.IsValid() {

		return fmt.Errorf("unknown field: %s\n", fieldName)
	}

	if !field.CanSet() || field.Kind() != reflect.String {
		return fmt.Errorf("Field %v cant be set as a string", fieldName)
	}

	field.SetString(strings.TrimSpace(string(after)))

	return nil
}

func main() {

	request := Request{}

	requestString := "POST /api/users HTTP/1.1\r\n" +
		"Host: example.com\r\n" +
		"Content-Type: application/json\r\n" +
		"Accept: application/json\r\n" +
		"Content-Length: 21\r\n" +
		"\r\n" +
		`{"name":"John"}`

	slice := []byte(requestString)

	index := 0

	for i := 0; i < len(slice); i++ {

		if i+1 < len(slice) {

			if slice[i] == 13 && slice[i+1] == 10 {

				err := parseString(slice[index:i], &request)
				if err != nil {
					fmt.Println("Error:", err)
				}
				index = i + 2
			}
		}

	}

	fmt.Println(request)

}
