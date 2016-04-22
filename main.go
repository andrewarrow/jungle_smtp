package main

import "net"
import "fmt"
import "bufio"
import "time"

type Client struct {
	conn    net.Conn
	address string
	time    int64
	bufin   *bufio.Reader
	bufout  *bufio.Writer
}

func handleClient(client *Client) {
	fmt.Println("here")
}

func main() {
	listener, _ := net.Listen("tcp", "0.0.0.0:25")

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(&Client{
			conn:    conn,
			address: conn.RemoteAddr().String(),
			time:    time.Now().Unix(),
			bufin:   bufio.NewReader(conn),
			bufout:  bufio.NewWriter(conn),
		})
	}
}
