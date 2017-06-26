package pingpong

import (
	"log"
	"net"
	"strings"
)

const (
	message       = "Ping"
	StopCharacter = "\r\n\r\n"
)

func SocketClient(ip, port *string) {
	addr := strings.Join([]string{*ip, *port}, ":")
	conn, err := net.Dial("tcp", addr)

	defer conn.Close()

	if err != nil {
		log.Fatalln(err)
	}

	conn.Write([]byte(message))
	conn.Write([]byte(StopCharacter))
	log.Printf("Send: %s", message)

	buff := make([]byte, 1024)
	n, _ := conn.Read(buff)
	log.Printf("Receive: %s", buff[:n])

}
