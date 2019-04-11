package listener

import (
	"crypto/tls"
	"fmt"
	"github.com/virepri/clipman-server/shared"
	"net"
	"os"
)

var listener net.Listener

func StartListener() {
	var err error

	if !shared.TLSEnabled {
		if listener, err = net.Listen("tcp", shared.BindTo); err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	} else {
		cert, err := tls.LoadX509KeyPair(shared.PubKeyPath, shared.PrivKeyPath)
		if err != nil {
			fmt.Println("Failed to load X509 keypair:", err.Error())
			os.Exit(0)
		}

		tlscfg := &tls.Config{Certificates: []tls.Certificate{cert}}
		if listener, err = tls.Listen("tcp", shared.BindTo, tlscfg); err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}

	for {
		if c, err := listener.Accept(); err == nil {
			d := shared.Device{Conn: c, HasAuth: false, HasAdmin: false}
			shared.Devices = append(shared.Devices, d)
			go handleConn(d)
		}
	}
}
