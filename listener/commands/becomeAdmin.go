package commands

import "github.com/virepri/clipman-server/shared"

func becomeAdmin(device *shared.Device, args []string) {
	if len(args) >= 1 {
		if args[0] == shared.PassHash {
			device.HasAdmin = true
		}
	}
}
