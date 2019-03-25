package commands

import (
	"github.com/virepri/clipman-server/shared"
	"github.com/virepri/clipman-server/user"
)

func login(device *shared.Device, args []string) {
	buffer := []byte{2, 0}
	if len(args) >= 1 {
		if args[0] == user.UserPassHash {
			device.HasAuth = true
			buffer[0] = 1
		}
	}

	_, _ = device.Conn.Write(buffer)
}
