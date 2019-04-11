package listener

import (
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"net"
	"os"
)

var listener net.Listener

//TODO: Add TLS
func StartListener() {
	var err error
	if listener, err = net.Listen("tcp", shared.BindTo); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}

	for {
		if c, err := listener.Accept(); err == nil {
			d := shared.Device{Conn: c, HasAuth: false, HasAdmin: false}
			shared.Devices = append(shared.Devices, d)
			go handleConn(d)
		}
	}
}
