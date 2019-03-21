package main

import (
	"fmt"
	"github.com/virepri/clipman-server/cli"
	"github.com/virepri/clipman-server/listener"
	"github.com/virepri/clipman-server/shared"
	"os/user"
)

func main() {
	shared.ActiveServices.Add(1)

	if u, err := user.Current(); err == nil {
		shared.ConfigLocation = u.HomeDir + "/.config/clipman-server.cfg"
	} else {
		fmt.Println("Can't get current user:")
		fmt.Println(err.Error())
		return
	}

	shared.LoadCFG()

	go cli.StartCLI()
	go listener.StartListener()

	shared.ActiveServices.Wait()
}
