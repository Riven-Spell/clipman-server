package shared

import (
	"encoding/json"
	"fmt"
	"github.com/virepri/clipman-server/user"
	"os"
	"path/filepath"
)

var BufferSize = 1024
var BindTo = ":7606"
var PassHash = "" //Hashes should be SHA256
var PrivKeyPath = ""
var PubKeyPath = ""
var TLSEnabled = false
var ConfigLocation = ""
var ConfigDir = ""
var PrintCMD = false

type cfg struct {
	Buffer         int
	Bind           string
	AdminHash      string
	UserHash       string
	PrivateKeyPath string
	PublicCertPath string
	TLSEnabled     bool
}

func SaveCFG() {
	if PrivKeyPath != "" {
		if pkp, err := filepath.Abs(PrivKeyPath); err == nil {
			PrivKeyPath = pkp
		}
	}

	if PubKeyPath != "" {
		if pkp, err := filepath.Abs(PubKeyPath); err == nil {
			PubKeyPath = pkp
		}
	}

	c := cfg{BufferSize, BindTo, PassHash, user.UserPassHash, PrivKeyPath, PubKeyPath, TLSEnabled}

	_ = os.MkdirAll(ConfigDir, 0600) //Who cares

	buff, _ := json.Marshal(c)
	if f, err := os.OpenFile(ConfigLocation, os.O_CREATE|os.O_RDWR, 0666); err == nil {
		if _, err := f.Write(buff); err != nil {
			fmt.Println(err.Error())
			return
		}
	} else {
		fmt.Println(err.Error())
	}
}

func LoadCFG() {
	var c cfg

	if fi, err := os.Stat(ConfigLocation); err == nil {
		buff := make([]byte, fi.Size())
		if f, err := os.Open(ConfigLocation); err == nil {
			if _, err := f.Read(buff); err == nil {
				if err := json.Unmarshal(buff, &c); err == nil {
					BufferSize = c.Buffer
					BindTo = c.Bind
					PassHash = c.AdminHash
					user.UserPassHash = c.UserHash
					PrivKeyPath = c.PrivateKeyPath
					PubKeyPath = c.PublicCertPath
					TLSEnabled = c.TLSEnabled
					return
				}
			}
		}
	}

	fmt.Println("Can't find a config. Creating a default config.")
	SaveCFG()
}
