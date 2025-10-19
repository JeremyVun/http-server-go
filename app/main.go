package main

import (
	"io"
	"fmt"
	"net"
	"os"
	"log"
)

// Ensures gofmt doesn't remove the "net" and "os" imports above (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", ":4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
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
	// Read from connection until we got everything
	request, err := parseRequest(conn)
	if (err != nil) {
		log.Fatal(err)
		return
	}

	response := handleRequest(request)

	// write to conn
	io.WriteString(conn, "HTTP/1.1 " + string(response.Status) + "\r\n" +
		"Content-Type: text/plain; charset=utf-8\r\n" +
		"Content-Length: 12\r\n" +
		"\r\n" +
		"hello world\n",
	)
	println("RESPONSE WRITTEN")
}
