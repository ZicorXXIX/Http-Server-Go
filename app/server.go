("fmt"
"net"
"os"
"path/filepath"
"strings"
)
func main() {
fmt.Println("Logs from your program will appear here!")
l, err := net.Listen("tcp", "0.0.0.0:4221")
if err != nil {
	fmt.Println("Failed to bind to port 4221")
	os.Exit(1)
}
for {
	conn, acceptErr := l.Accept()
	if acceptErr != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	go func() {
		defer conn.Close()
		reply := make([]byte, 1024)
		_, err = conn.Read(reply)
		if err != nil {
			println("Write to server failed:", err.Error())
			os.Exit(1)
		}
		firstLine := strings.Split(string(reply), "\r\n")[0]
		reqPath := strings.Split(firstLine, " ")[1]
		if strings.Split(reqPath, "/")[1] == "echo" {
			value := strings.Split(reqPath, "/echo/")[1]
			resp := fmt.Sprintf("HTTP/1.1 %d %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s\r\n", 200, "OK", len(value), value)
			conn.Write([]byte(resp))
		} else if strings.Split(reqPath, "/")[1] == "files" {
			value := strings.Split(reqPath, "/files/")[1]
			directory := os.Args[2]
			dat, err := os.ReadFile(filepath.Join(directory, value))
			fmt.Println(filepath.Join(directory, value), err)
			if err != nil {
				returnedBytes := []byte("HTTP/1.1 404 Not Found\r\n\r\n")
				conn.Write(returnedBytes)
			}
			resp := fmt.Sprintf("HTTP/1.1 %d %s\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s\r\n", 200, "OK", len(dat), string(dat))
1
			conn.Write([]byte(resp))
		} else if strings.Split(reqPath, "/")[1] == "user-agent" {
			lines := strings.Split(string(reply), "\r\n")
			id := 0
			for i := 0; i < len(lines); i++ {
				if strings.Contains(lines[i], "User-Agent") {
					id = i
				}
			}
			fmt.Println(lines[id], id)
			value := strings.Split(lines[id], " ")[1]
			resp := fmt.Sprintf("HTTP/1.1 %d %s\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s\r\n", 200, "OK", len(value), value)
			conn.Write([]byte(resp))
		} else if strings.Split(firstLine, " ")[1] != "/" {
			returnedBytes := []byte("HTTP/1.1 404 Not Found\r\n\r\n")
			conn.Write(returnedBytes)
		} else {
			returnedBytes := []byte("HTTP/1.1 200 OK\r\n\r\n")
			conn.Write(returnedBytes)
		}
	}()
}
}