package main

type HttpStatusCode string
const (
	OK HttpStatusCode = "200 OK"
	NOT_FOUND HttpStatusCode = "404 Not Found"
)

type HttpResponse struct {
	Status HttpStatusCode
	Body string
}
