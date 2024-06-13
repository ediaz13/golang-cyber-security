package main

import (
	"fmt"
	"net"
)

const (
    targetHost = "127.0.0.1"
    targetPort = "9998"
)

func main() {
    // Create a connection to the server
    conn, err := net.Dial("tcp", targetHost+":"+targetPort)
    if err != nil {
        fmt.Println("Error connecting to server:", err)
        return
    }
    defer conn.Close()

    // Send some data
    request := "GET / HTTP/1.1 \r\nHost: hacker.com\r\n\r\n"
    _, err = conn.Write([]byte(request))
    if err != nil {
        fmt.Println("Error sending data:", err)
        return
    }

    // Receive some data
    buffer := make([]byte, 4096)
    n, err := conn.Read(buffer)
    if err != nil {
        fmt.Println("Error receiving data:", err)
        return
    }

    fmt.Println(string(buffer[:n]))
}