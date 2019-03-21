package commands

import "github.com/virepri/clipman-server/shared"

func becomeAdmin(args []string, hasAuth *bool, hasAdmin *bool) {
	if len(args) >= 1 {
		if args[0] == shared.PassHash {
			*hasAdmin = true
		}
	}
}
