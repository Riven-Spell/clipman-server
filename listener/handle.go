package listener

import (
	"github.com/virepri/clipman-server/listener/commands"
	"github.com/virepri/clipman-server/shared"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()

	buff := make([]byte, shared.BufferSize)
	for {
		if _, err := conn.Read(buff); err == nil {
			cmd := commands.ParseCmd(buff)

			if v, ok := commands.Aliases[cmd.Cmd]; ok {
				v(cmd.Args)
			}
		}
	}
}
