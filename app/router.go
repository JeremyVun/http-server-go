package main

import(
	"fmt"
)

var router = map[RouteKey]func(*HttpRequest)HttpResponse {
	{Method: GET, Path: "/path"}: Hello,
}

func Hello(request *HttpRequest) HttpResponse {
	println("hello!")
	return HttpResponse {
		Status: OK,
	}
}

func handleRequest(request *HttpRequest) HttpResponse {
	fmt.Printf("%+v\n", request)
	if handler, exists := router[request.RouteKey]; exists {
		return handler(request)
	}

	return HttpResponse {
		Status: NOT_FOUND,
	}
}
