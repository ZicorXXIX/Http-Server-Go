package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	// Uncomment this block to pass the first stage
	//
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	//
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	
	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	statusHeader := strings.Split(string(buf),"\r\n")[0]
	route := strings.Split(statusHeader, " ")[1]
	
	if route == "/" {
		sendResponse(conn, "HTTP/1.1 200 OK\r\n\r\nHello, World!")
		} else {
			sendResponse(conn, "HTTP/1.1 404 Not Found\r\n\r\n404 Not Found")
			
	}
	// response := "HTTP/1.1 200 OK\r\n\r\n"
	// _, err = conn.Write([]byte(response))
	
}


func sendResponse(conn net.Conn, response string) {
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
	}
}
