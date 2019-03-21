package commands

import (
	"github.com/virepri/clipman-server/shared"
	"github.com/virepri/clipman-server/user"
)

func getClip(device *shared.Device, args []string) {
	if device.HasAuth {
		buff := []byte{0, 10}
		buff = append(buff, []byte(user.ClipboardContent)...)

		_,_ = device.Conn.Write(buff)
	}
}
