package main

var router = map[RouteKey]func(*HttpRequest)HttpResponse {
	{Method: GET, Path: "/hello"}: Hello,
}

func Hello(request *HttpRequest) HttpResponse {
	return HttpResponse {
		Status: OK,
		Body: "hello world!",
	}
}

func handleRequest(request *HttpRequest) HttpResponse {
	if handler, exists := router[request.RouteKey]; exists {
		return handler(request)
	}

	return HttpResponse {
		Status: NOT_FOUND,
	}
}
