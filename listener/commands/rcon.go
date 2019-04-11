package commands

import (
	"github.com/virepri/clipman-server/cli/commands"
	"github.com/virepri/clipman-server/shared"
)

func rcon(device *shared.Device, args []string) {
	if device.HasAdmin {
		if len(args) >= 1 {
			if v, ok := commands.Aliases[args[0]]; ok {
				go v(args[1:])
			}
		}
	}
}
