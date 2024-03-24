package host

import (
	"net"
	"sync"
)

type Host struct {
	Domain string
	Port   uint16
	Conn   net.Conn
	mu     sync.Mutex
}
