package commands

import (
	"github.com/virepri/clipman-server/shared"
	"github.com/virepri/clipman-server/user"
)

func setClip(device *shared.Device, args []string) {
	if device.HasAuth && len(args) >= 1 {
		user.ClipboardContent = args[0]
		buff := []byte{0, 10}
		buff = append(buff, []byte(user.ClipboardContent)...)

		for _,d := range shared.Devices {
			_,_ = d.Conn.Write(buff)
		}
	}
}
