package main

import (
    "fmt"
    "net"
    "os"
)

const (
    CONN_HOST = "localhost"
    CONN_PORT = "13333"
    CONN_TYPE = "tcp"
)

func main() {
    // listen for incoming connections
    l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

    // close the listener when the application closes
    defer l.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    // listen for incoming connections
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println(err.Error())
            os.Exit(1)
        }
	// send request to our handler
        go handleRequest(conn)
    }
}

// handles incoming request
func handleRequest(conn net.Conn) {
  // make a buffer to hold incoming data
  buf := make([]byte, 1024)
  // Read the incoming connection into the buffer.
  _, err := conn.Read(buf)
  if err != nil {
    fmt.Println("Error reading:", err.Error())
  }
  // Send a response back to person contacting us.
  conn.Write([]byte("Message received."))
  // Close the connection when you're done with it.
  conn.Close()
}
