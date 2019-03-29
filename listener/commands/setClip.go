package commands

import (
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"github.com/virepri/clipman-server/user"
	"strings"
)

func setClip(device *shared.Device, args []string) {
	if device.HasAuth && len(args) >= 1 {
		user.ClipboardContent = strings.Join(args, "\n")

		buff := []byte{0, 10}
		buff = append(buff, []byte(user.ClipboardContent)...)
		buff = append(buff, 0)

		fmt.Println(buff)

		for _,d := range shared.Devices {
			_,_ = d.Conn.Write(buff)
			fmt.Println(d, buff)
		}
	}
}
