package commands

import "fmt"

func help(args []string) {
	if len(args) >= 1 {
		for k,v := range args {
			fmt.Println(k, v)
		}

		if v, ok := HelpList[args[0]]; ok {
			fmt.Println(args[0] + ":\n" + v + "\n")
		} else {
			fmt.Println("No help exists for command " + args[0])
		}
	} else {
		for k, v := range HelpList {
			fmt.Println(k + ":\n" + v + "\n")
		}
	}
}
