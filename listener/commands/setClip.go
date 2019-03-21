package commands

import (
	"github.com/virepri/clipman-server/shared"
	"github.com/virepri/clipman-server/user"
)

func setClip(device *shared.Device, args []string) {
	if device.HasAuth && len(args) >= 1 {
		user.ClipboardContent = args[0]
		//TODO: divvy out the new content to clients.
	}
}
