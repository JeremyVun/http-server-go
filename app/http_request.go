package main

type RequestLine struct {
	Scheme string
	Method HttpMethod
	Path string
}

type HttpRequest struct {
	RequestLine
	Headers map[string]string
	Body *string
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
