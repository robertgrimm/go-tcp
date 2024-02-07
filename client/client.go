package main

import (
	"fmt"
	"net"
	"strconv"
)

const (
	PORT = 9000
)

func main() {
	fmt.Println("starting tcp client")

	c, err := net.Dial("tcp", "localhost:"+strconv.Itoa(PORT))
	if err != nil {
		fmt.Println("dial failed:", err)
		return
	}

	msg := "Hello, Server!\n"
	_, err = c.Write([]byte(msg))
	if err != nil {
		fmt.Println("write failed:", err)
		return
	}

	_ = c.Close()

	fmt.Println("tcp client stopped")
}
