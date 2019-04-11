package shared

import (
	"encoding/json"
	"fmt"
	"github.com/virepri/clipman-server/user"
	"os"
)

var BufferSize = 1024
var BindTo = ":7606"
var PassHash = "" //Hashes should be SHA256
var ConfigLocation = ""
var PrintCMD = false

type cfg struct {
	Buffer    int
	Bind      string
	AdminHash string
	UserHash string
}

func SaveCFG() {
	c := cfg{BufferSize, BindTo, PassHash, user.UserPassHash}

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
					return
				}
			}
		}
	}

	fmt.Println("Can't find a config. Creating a default config.")
	SaveCFG()
}
