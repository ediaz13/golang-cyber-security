# TCP Server and Client for Cybersecurity in Go

## Overview

This repository contains a simple TCP server and client written in Go, designed for educational purposes in the context of cybersecurity. The server listens for incoming connections and responds with a basic HTTP response. The client connects to the server, sends a HTTP GET request, and prints the server's response.

## Contents

- **server.go**: The Go implementation of the TCP server.
- **client.go**: The Go implementation of the TCP client.

## Requirements

- Go 1.14 or later

## Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/tcp-server-client-go.git
    cd tcp-server-client-go
    ```

2. **Ensure Go is installed:**

    Make sure you have Go installed on your machine. You can download it from the official [Go website](https://golang.org/dl/).

## Usage

### Running the Server

1. **Navigate to the directory:**

    ```sh
    cd path/to/your/repository
    ```

2. **Compile and run the server:**

    ```sh
    go run server.go
    ```

3. The server will start listening on `0.0.0.0:9998`.

### Running the Client

1. **Open a new terminal window or tab.**

2. **Navigate to the directory:**

    ```sh
    cd path/to/your/repository
    ```

3. **Compile and run the client:**

    ```sh
    go run client.go
    ```

4. The client will connect to `127.0.0.1:9998`, send a HTTP GET request, and print the response from the server.

## Code Explanation

### server.go

```go
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
    if (err != nil) {
        fmt.Println("Error creating server:", err)
        return
    }
    defer listener.Close()
    fmt.Printf("[*] Listening on %s:%s\n", IP, PORT)

    for {
        conn, err := listener.Accept()
        if (err != nil) {
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
    if (err != nil) {
        fmt.Println("Error reading request:", err)
        return
    }
    fmt.Printf("[*] Received: %s", strings.TrimSpace(request))
    response := "HTTP/1.1 200 OK\r\n"
    conn.Write([]byte(response))
}
