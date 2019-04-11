package commands

import (
	"fmt"
	"github.com/virepri/clipman-server/shared"
	user2 "github.com/virepri/clipman-server/user"
)

func debug(args []string) {
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "clipboard":
		fmt.Println(user2.ClipboardContent)
	case "devices":
		for _,v := range shared.Devices {
			fmt.Println(v)
		}
	case "printcmd":
		shared.PrintCMD = !shared.PrintCMD
		fmt.Println("Toggled PrintCMD:", shared.PrintCMD)
	}
}
