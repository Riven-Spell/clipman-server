package listener

import (
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"net"
	"os"
)

var listener net.Listener

func StartListener() {
	var err error
	if listener, err = net.Listen("tcp", shared.BindTo); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	if c, err := listener.Accept(); err == nil {
		go handleConn(c)
	}
}