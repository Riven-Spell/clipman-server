package commands

import "github.com/virepri/clipman-server/cli/commands"

func rcon(args []string, hasAuth *bool, hasAdmin *bool) {
	if *hasAdmin {
		if len(args) >= 1 {
			if v, ok := commands.Aliases[args[0]]; ok {
				go v(args[1:])
			}
		}
	}
}