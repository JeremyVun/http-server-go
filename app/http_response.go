package main

import(
	"bytes"
	"fmt"
	"strconv"
)

type HttpStatusCode string
const (
	OK HttpStatusCode = "200 OK"
	NOT_FOUND HttpStatusCode = "404 Not Found"
)

type HttpResponse struct {
	Scheme string
	Status HttpStatusCode
	Headers map[string]string
	Body *string
}

func (r HttpResponse) ToHttpResponseString(isPipeliningAllowed bool) string {
	var buffer bytes.Buffer

	// response line e.g. HTTP/1.1 200
	buffer.WriteString(fmt.Sprintf("%s %s\r\n", r.Scheme, r.Status))

	// Headers
	if (isPipeliningAllowed) {
		r.Headers["Connection"] = "keep-alive"
	} else {
		r.Headers["Connection"] = "close"
	}
	if (r.Body == nil) {
		r.Headers["Content-Length"] = "0"
	} else {
		r.Headers["Content-Length"] = strconv.Itoa(len(*r.Body))
	}

	for k, v := range(r.Headers) {
		buffer.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	buffer.WriteString("\r\n")

	// Body
	if (r.Body != nil) {
		buffer.WriteString(*r.Body)
	}

	return buffer.String()
}
