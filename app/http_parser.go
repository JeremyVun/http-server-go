package main

import (
	"io"
	"bufio"
	"strings"
	"log"
)

// Parse the request line
// e.g. GET /path HTTP/1.1\r\n
func parseRequestLine(reader *bufio.Reader) (*RequestLine, error) {
	bytes, err := reader.ReadBytes('\n')
	if (err != nil) {
		return nil, err
	}

	requestLine := strings.Split(string(bytes), " ")

	return &RequestLine {
		Method: HttpMethod(requestLine[0]),
		Scheme: strings.TrimRight(requestLine[2], "\r\n"),
		Path: requestLine[1],
	}, nil
}

func parseHeaders(reader *bufio.Reader) map[string]string  {
	headers := make(map[string]string)

	for {
		bytes, _ := reader.ReadBytes('\n')
		headerLine := strings.ToLower(string(bytes))
		if (headerLine == "\r\n") {
			break
		}
		delimiterIndex := strings.IndexByte(headerLine, ':')
		if (delimiterIndex >= 0) {
			headers[headerLine] = ""
		} else {
			headers[headerLine[:delimiterIndex]] = headerLine[delimiterIndex:]
		}
	}

	return headers
}

func parseBody(reader *bufio.Reader, contentLength int) *string {
	if contentLength <= 0 {
		return nil
	}

	buffer := make([]byte, contentLength)
	n, err := io.ReadFull(reader, buffer)
	if (n != contentLength || err != nil) {
		log.Printf("Cannot read body with content-length %d - error: %s", contentLength, err)
	}

	bodyString := string(buffer)
	return &bodyString
}
