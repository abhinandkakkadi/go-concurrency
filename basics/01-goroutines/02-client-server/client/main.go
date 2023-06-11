package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	
	defer conn.Close()

	mustCopy(os.Stdout,conn)

}


func mustCopy(dst io.Writer,src io.Reader) {

	_,err := io.Copy(dst,src)
	if err != nil {
		fmt.Println(err)
	}
}