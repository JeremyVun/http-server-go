package main

import (
	"strconv"
	"io"
	"fmt"
	"net"
	"os"
	"log"
	"flag"
)

func main() {
	portPtr := flag.Int("p", 8080, "Port")
	flag.Parse()
	port := *portPtr
	fmt.Printf("Starting server on port %d\n", port)

	l, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n - %s", port, err)
		os.Exit(1)
	}
	fmt.Printf("Listening on port %d\n", port)
	defer l.Close() // close the listener when this function ends

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Parse the request from the Http Connection
	request, err := parseRequest(conn)
	if (err != nil) {
		log.Fatal(err)
		return
	}

	response := handleRequest(request)

	// Write response back to connection
	contentLength := len(response.Body) +1
	io.WriteString(conn, "HTTP/1.1 " + string(response.Status) + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"Content-Length: " + strconv.Itoa(contentLength) + "\r\n" +
		"\r\n" +
		response.Body + "\n",
	)
}
