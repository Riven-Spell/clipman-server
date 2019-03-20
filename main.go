package main

import (
	"github.com/virepri/clipman-server/cli"
	"github.com/virepri/clipman-server/listener"
	"github.com/virepri/clipman-server/shared"
)

func main() {
	shared.ActiveServices.Add(1)

	go cli.StartCLI()
	go listener.StartListener()

	shared.ActiveServices.Wait()
}
