package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	// fmt.Println("System Args:", os.Args[1])
	
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	//
	for{
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		go handleConnection(conn)
	}
	
	
	// response := "HTTP/1.1 200 OK\r\n\r\n"
	// _, err = conn.Write([]byte(response))
	
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// conn.SetReadDeadline(time.Now().Add(1000 * time.Second))
	// conn.SetWriteDeadline(time.Now().Add(1000 * time.Second))
	// conn.SetReadDeadline(time.Now().Add(100 * time.Second))
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
			sendResponse(conn, "HTTP/1.1 404 Not Found\r\n\r\n404 Not Found")
		} else if route == "/user-agent" {
			// parseUserAgent(conn)
			// fmt.Println(userAgent)
			userAgent := strings.Split(headers[2], " ")[1] 
			response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length:%d\r\n\r\n%s", len(userAgent), userAgent)
			sendResponse(conn, response) 
		} else if strings.Contains(route, "/files"){
			fileName := strings.Split(route, "/files/")[1]
			filepath := os.Args[2] + fileName
			fmt.Println("FILE PATH:",filepath)
			f, err := os.ReadFile(filepath)
			fmt.Println("FILE CONTENTS:",string(f))
			if err != nil {
				sendResponse(conn, "HTTP/1.1 404 Not Found\r\n\r\n")						
			} else {
				data := string(f)
				response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length:%d\r\n\r\n%s", len(f), data)
				sendResponse(conn, response)
			}

			
		} else {
			sendResponse(conn, "HTTP/1.1 404 Not Found\r\n\r\n404 Not Found")
			
	}
}


func sendResponse(conn net.Conn, response string) {
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response: ", err.Error())
		conn.Close()
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
