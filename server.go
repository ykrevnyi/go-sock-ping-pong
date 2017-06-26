package pingpong

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

const (
	Message = "Pong"
)

func SocketServer(port *string) {

	listen, err := net.Listen("tcp4", ":"+*port)
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %s failed,%s", *port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %s", *port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	defer conn.Close()

	var (
		buf = make([]byte, 1024)
		r   = bufio.NewReader(conn)
		w   = bufio.NewWriter(conn)
	)

ILOOP:
	for {
		n, err := r.Read(buf)
		data := string(buf[:n])

		switch err {
		case io.EOF:
			break ILOOP
		case nil:
			log.Println("Receive:", data)
			if isTransportOver(data) {
				break ILOOP
			}

		default:
			log.Fatalf("Receive data failed:%s", err)
			return
		}

	}
	w.Write([]byte(Message))
	w.Flush()
	log.Printf("Send: %s", Message)

}

func isTransportOver(data string) (over bool) {
	over = strings.HasSuffix(data, "\r\n\r\n")
	return
}
