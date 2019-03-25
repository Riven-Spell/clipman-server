package commands

import "github.com/virepri/clipman-server/shared"

func becomeAdmin(device *shared.Device, args []string) {
	buffer := []byte{2, 0}
	if len(args) >= 1 && device.HasAuth {
		if args[0] == shared.PassHash {
			device.HasAdmin = true
			buffer[0] = 1
		}
	}

	_, _ = device.Conn.Write(buffer)
}
