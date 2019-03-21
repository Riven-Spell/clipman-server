package commands

import "fmt"

func help(args []string) {
	for k, v := range HelpList {
		fmt.Println(k + ":\n" + v + "\n")
	}
}
