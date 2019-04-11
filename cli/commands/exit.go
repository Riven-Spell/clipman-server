package commands

import (
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"os"
)

func exit(args []string) {
	fmt.Println("Shutting down the server.")
	shared.ActiveServices.Done()
	os.Exit(0)
}
