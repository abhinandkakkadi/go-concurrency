package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

func main() {

	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}

	for {

		conn, err := ln.Accept()
		if err != nil {
			continue
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, "connection from server")
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
