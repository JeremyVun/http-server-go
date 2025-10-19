package main

import (
	"bufio"
	"strings"
	"net"
)

type RouteKey struct {
	Method HttpMethod
	Path string
}

type HttpRequest struct {
	RouteKey
	Scheme string
	Headers map[string]string
}

type HttpMethod string
const (
	GET HttpMethod = "GET"
	POST HttpMethod = "POST"
	PUT HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
	PATCH HttpMethod = "PATCH"
)
var validMethods = map[HttpMethod]struct{} {
	GET: {},
	POST: {},
	PUT: {},
	DELETE: {},
	PATCH: {},
}

func parseRequest(conn net.Conn) (*HttpRequest, error) {
	r := bufio.NewReader(conn)

	// parse request line
	requestLineBytes, _ := r.ReadBytes('\n')
	requestLineTokens := strings.Split(string(requestLineBytes), " ")
	method := HttpMethod(requestLineTokens[0])
	path := requestLineTokens[1]
	scheme := requestLineTokens[2]

	// parse headers
	headers := make(map[string]string)
	for {
		bytes, _ := r.ReadBytes('\n')
		line := string(bytes)
		if (line == "\r\n") {
			break
		}
		header := strings.SplitN(line, ":", 2)
		headers[header[0]] = header[1]
	}

	// parse body

	return &HttpRequest {
		RouteKey: RouteKey{Method: method, Path: path},
		Scheme: scheme,
		Headers: headers,
	}, nil
}
