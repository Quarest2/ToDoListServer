package host

import (
	log "github.com/sirupsen/logrus"
	"net"
	"sync"
)

var host Host

func Init() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error with listening the mainServer.")
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Println("Error with accepting the mainServer.")
	}
	host = Host{Domain: "localhost", Port: 8080, Conn: conn, mu: sync.Mutex{}}
}
