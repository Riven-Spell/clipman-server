package listener

import (
	"github.com/virepri/clipman-server/shared"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, shared.BufferSize)
	for {
		if _, err := conn.Read(buff); err == nil {
			//TODO: handle incoming commands.
		}
	}
}
