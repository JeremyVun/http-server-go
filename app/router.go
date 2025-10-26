package main

import (
	"fmt"
	"time"
)

type RouteKey struct {
	Method HttpMethod
	Path string
}

var router = map[RouteKey]func(*HttpRequest)HttpResponse {
	{Method: GET, Path: "/hello"}: HelloHandler,
	{Method: GET, Path: "/expensive"}: ExpensiveHandler,
}

func HelloHandler(request *HttpRequest) HttpResponse {
	body := "hello world!"
	return HttpResponse {
		Scheme: request.Scheme,
		Status: OK,
		Body: &body,
		Headers: make(map[string]string),
	}
}

func ExpensiveHandler(request *HttpRequest) HttpResponse {
	fmt.Println("Starting expensive request")

	time.Sleep(1 * time.Second)

	body := "expensive request!"
	return HttpResponse {
		Scheme: request.Scheme,
		Status: OK,
		Body: &body,
		Headers: make(map[string]string),
	}
}

func handleRequest(request *HttpRequest) HttpResponse {
	routeKey := RouteKey {
		Method: request.RequestLine.Method,
		Path: request.RequestLine.Path,
	}
	if handler, exists := router[routeKey]; exists {
		return handler(request)
	}

	return HttpResponse {
		Scheme: request.Scheme,
		Status: NOT_FOUND,
		Headers: make(map[string]string),
	}
}
