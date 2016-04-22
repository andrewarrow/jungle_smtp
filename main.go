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

func (c *Client) w(s string) {
	c.bufout.WriteString(s + "\r\n")
	c.bufout.Flush()
}
func (c *Client) r() string {
	reply, err := c.bufin.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	return reply
}

func handleClient(c *Client) {
	fmt.Println("here")
	c.w("220 Welcome to the Jungle")
	text := c.r()
	c.w("250 No one says helo anymore")
	text = c.r()
	c.w("250 Sender")
	text = c.r()
	c.w("250 Recipient")
	text = c.r()
	c.w("354 Ok Send data ending with <CRLF>.<CRLF>")
	text = c.r()
	fmt.Println(text)
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
