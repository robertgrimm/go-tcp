package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	PORT = 9000
)

func main() {
	fmt.Println("starting tcp server")

	l, err := net.Listen("tcp", ":"+strconv.Itoa(PORT))
	if err != nil {
		fmt.Println("net listener failed:", err)
		return
	}
	defer func() { _ = l.Close() }()

	waitForEnter()

	for {
		var c net.Conn
		c, err = l.Accept()
		if err != nil {
			fmt.Println("accept failed:", err)
			return
		}
		go handleConnection(c)
	}
}

func waitForEnter() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func handleConnection(c net.Conn) {
	defer func() { _ = c.Close() }()

	fmt.Println("handling connection")

	buf := make([]byte, 1024)
	n, err := c.Read(buf)
	if err != nil {
		fmt.Println("read failed:", err)
		return
	}

	fmt.Println("read", n, "bytes:", string(buf[:n]))

	t := time.Now()
	res := t.Format(time.RFC3339) + "\n"
	_, err = c.Write([]byte(res))
	if err != nil {
		fmt.Println("failed to write response:", err)
		return
	}
}
