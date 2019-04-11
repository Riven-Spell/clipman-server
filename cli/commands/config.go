package commands

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"strconv"
)

func config(args []string) {
	if len(args) >= 1 {
		switch args[0] {
		case "save":
			shared.SaveCFG()
			fmt.Println("Wrote config.")
		case "reload":
			shared.LoadCFG()
			fmt.Println("Reloaded config.")
		case "buffer":
			if len(args) >= 2 {
				if n, err := strconv.Atoi(args[1]); err == nil {
					shared.BufferSize = n
				} else {
					fmt.Println("Buffer is an integer. Please enter an integer.")
				}
			}
			fmt.Println("The buffer size is", shared.BufferSize)
		case "bind":
			if len(args) >= 2 {
				shared.BindTo = args[1]
			}
			fmt.Println("The tcp listener binds to", shared.BindTo)
		case "password":
			if len(args) >= 2 {
				h := sha256.New()
				h.Write([]byte(args[1]))
				shared.PassHash = hex.EncodeToString(h.Sum(nil))
				fmt.Println("Your password has been changed.")
			} else {
				fmt.Println("Please supply a password.")
			}
		case "tls":
			if len(args) >= 2 {
				switch args[1] {
				case "certlocation":
					if len(args) > 3 {
						shared.PubKeyPath = args[2]
					}
					fmt.Println("Certificate is stored @", shared.PubKeyPath)
				case "keylocation":
					if len(args) > 3 {
						shared.PrivKeyPath = args[2]
					}
					fmt.Println("Private key is stored @", shared.PrivKeyPath)
				case "toggle":
					shared.TLSEnabled = !shared.TLSEnabled
					fmt.Println("TLS enabled:", shared.TLSEnabled)
				default:
					fmt.Println("Certificate is stored @", shared.PubKeyPath)
					fmt.Println("Private key is stored @", shared.PrivKeyPath)
					fmt.Println("TLS enabled:", shared.TLSEnabled)
				}
			}
		default:
			help([]string{"config"})
		}
	} else {
		fmt.Println("Config info:")
		fmt.Println("Buffer size", shared.BufferSize)
		fmt.Println("TCP listener is binding to", shared.BindTo)
	}
}
