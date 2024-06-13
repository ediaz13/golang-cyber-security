package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

const (
    IP   = "0.0.0.0"
    PORT = "9998"
)

func main() {
    listener, err := net.Listen("tcp", IP+":"+PORT)
    if err != nil {
        fmt.Println("Error creating server:", err)
        return
    }
    defer listener.Close()
    fmt.Printf("[*] Listening on %s:%s\n", IP, PORT)

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        fmt.Printf("[*] Accepted connection from %s\n", conn.RemoteAddr().String())
        go handleClient(conn)
    }
}

func handleClient(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    request, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Error reading request:", err)
        return
    }
    fmt.Printf("[*] Received: %s", strings.TrimSpace(request))
    response := "HTTP/1.1 200 OK\r\n"
    conn.Write([]byte(response))
}