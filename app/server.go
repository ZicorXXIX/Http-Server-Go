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
	
	// route := parseRoute(conn)
	headers := parseHeaders(conn)
	route := strings.Split(headers[0], " ")[1]
	// userAgent := parseHeaders(conn)
	
	if route == "/" {
		sendResponse(conn, "HTTP/1.1 200 OK\r\n\r\nHello, World!")
		} else if strings.Contains(route, "/echo"){
			responseBody := strings.Split(route, "echo/")[1]
			fmt.Println(responseBody)
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%d\r\n\r\n%s", len(responseBody), responseBody)
			sendResponse(conn, response)
		} else if route == "/user-agent" {
			// parseUserAgent(conn)
			// fmt.Println(userAgent)
			userAgent := strings.Split(headers[2], " ")[1] 
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%d\r\n\r\n%s", len(userAgent), userAgent)
			sendResponse(conn, response) 
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

// func parseRoute(conn net.Conn) string {
// 	buf := make([]byte, 1024)
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading request: ", err.Error())
// 	}
// 	statusHeader := strings.Split(string(buf),"\r\n")[0]
// 	route := strings.Split(statusHeader, " ")[1]
// 	return route		
// }

func parseHeaders(conn net.Conn) []string {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading request: ", err.Error())
	}
	statusHeader := strings.Split(string(buf),"\r\n")
	return statusHeader
}
