package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/virepri/clipman-server/shared"
	user2 "github.com/virepri/clipman-server/user"
)

func user(args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "save":
			shared.SaveCFG()
		case "password":
			if len(args) >= 2 {
				hash := sha256.New()
				hash.Write([]byte(args[1]))
				user2.UserPassHash = hex.EncodeToString(hash.Sum(nil))
				fmt.Println("Changed the user's password.")
			}
		}
	}
}
