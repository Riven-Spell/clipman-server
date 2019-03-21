package listener

import (
	"github.com/virepri/clipman-server/listener/commands"
	"github.com/virepri/clipman-server/shared"
)

func handleConn(device shared.Device) {
	defer device.Conn.Close()

	buff := make([]byte, shared.BufferSize)
	for {
		if _, err := device.Conn.Read(buff); err == nil {
			cmd := commands.ParseCmd(buff)

			if v, ok := commands.Aliases[cmd.Cmd]; ok {
				go v(&device, cmd.Args)
			}
		}
	}
}
