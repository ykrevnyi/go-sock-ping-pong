package main

import (
	"flag"

	pingpong "github.com/ykrevnyi/go-sock-ping-pong"
)

func main() {
	var (
		ip   = flag.String("ip", "127.0.0.1", "IP address")
		port = flag.String("port", "8000", "HTTP listen port")
	)
	flag.Parse()

	pingpong.SocketClient(ip, port)
}
