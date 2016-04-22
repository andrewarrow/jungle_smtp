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

func handleClient(c *Client) {
	fmt.Println("here")
	c.bufout.WriteString("220 Welcome to the Jungle\r\n")
	c.bufout.Flush()
	reply, err := c.bufin.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(reply)
	c.bufout.WriteString("250 No one says helo anymore.\r\n")
	c.bufout.Flush()
}

func main() {
	listener, _ := net.Listen("tcp", "0.0.0.0:2525")

	for {
		fmt.Println("waiting...")
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
