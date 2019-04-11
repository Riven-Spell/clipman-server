package shared

import "net"

type Device struct {
	Conn     net.Conn
	HasAuth  bool
	HasAdmin bool
}
