package commands

type Command func(args []string)
var Aliases = map[string]Command {
	"help":help,
}

var HelpList = map[string]string {
	"help":`displays this text`,
}