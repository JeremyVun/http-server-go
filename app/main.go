package main

import (
	"bufio"
	"strconv"
	"io"
	"fmt"
	"net"
	"os"
	"flag"
	"time"
)

func main() {
	portPtr := flag.Int("p", 8080, "Port")
	keepAlivePtr := flag.Bool("ka", false, "Keep Alive")
	flag.Parse()
	port := *portPtr
	isPipeliningAllowed := *keepAlivePtr
	fmt.Printf("Starting server on port %d. Pipelining: %v\n", port, isPipeliningAllowed)

	l, err := net.Listen("tcp", ":" + strconv.Itoa(port))
	if err != nil {
		fmt.Printf("Failed to bind to port %d\n - %s", port, err)
		os.Exit(1)
	}

	fmt.Printf("Listening on port %d\n", port)
	defer l.Close()

	// http server event loop
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn, isPipeliningAllowed)
	}
}

func handleConnection(conn net.Conn, isPipeliningAllowed bool) {
	// fmt.Println("Handling new connection from", conn.RemoteAddr())
	conn.SetDeadline(time.Now().Add(10 * time.Second))
	defer conn.Close()

	for {
		// Parse request from the Connection
		reader := bufio.NewReader(conn)
		requestLine, err := parseRequestLine(reader)
		if (err != nil) {
			break
		}

		headers := parseHeaders(reader)
		requestContentLength, err := strconv.Atoi(headers["content-length"])
		if (err != nil) {
			requestContentLength = -1
		}
		body := parseBody(reader, requestContentLength)

		request := HttpRequest {
			RequestLine: *requestLine,
			Headers: headers,
			Body: body,
		}

		// Handle the request
		response := handleRequest(&request)

		// Write response back to connection\
		io.WriteString(conn, response.ToHttpResponseString(isPipeliningAllowed))

		if (!isPipeliningAllowed) {
			break
		}
	}

	// fmt.Println("Closing connection from", conn.RemoteAddr())
}
