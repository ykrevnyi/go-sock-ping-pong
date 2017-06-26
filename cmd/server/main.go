package main

import (
	"flag"

	pingpong "github.com/ykrevnyi/go-sock-ping-pong"
)

func main() {
	var (
		port = flag.String("port", "8000", "HTTP listen port")
	)
	flag.Parse()

	// errs := make(chan error)
	// go func() {
	// 	c := make(chan os.Signal)
	// 	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// 	errs <- fmt.Errorf("%s", <-c)
	// 	os.Exit(0)
	// }()

	pingpong.SocketServer(port)
}
