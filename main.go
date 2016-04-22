package main

import "net"
import "fmt"
import "bufio"
import "time"
import "os"
import "io"

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
		fmt.Println("e ", err)
	}
	return reply
}

func appendToFile(text string) {
	of, _ := os.OpenFile("incoming_emails.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	defer of.Close()
	of.Seek(0, os.SEEK_END)
	io.WriteString(of, text+"\n")
}

func handleClient(c *Client) {
	fmt.Println("here")
	c.w("220 Welcome to the Jungle")
	text := c.r()
	appendToFile(text)
	c.w("250 No one says helo anymore")
	text = c.r()
	appendToFile(text)
	c.w("250 Sender")
	text = c.r()
	appendToFile(text)
	c.w("250 Recipient")
	text = c.r()
	appendToFile(text)
	c.w("354 Ok Send data ending with <CRLF>.<CRLF>")
	for {
		text = c.r()
		bytes := []byte(text)
		appendToFile(text)
		// 46 13 10
		if bytes[0] == 46 && bytes[1] == 13 && bytes[2] == 10 {
			break
		}
	}
	c.conn.Close()
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
