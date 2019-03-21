package commands

type Command func(args []string)

var Aliases = map[string]Command{
	"help":   help,
	"config": config,
	"exit": exit,
	"user": user,
}

var HelpList = map[string]string{
	"help": `displays this text`,
	"config": `config save:
saves the config.

config reload:
reloads the config, to undo changes. This is not equivalent to a server restart. 

config buffer [size]:
changes the buffer size.

config bind [port]:
changes where the server listens on.

config password [password]:
changes the admin password stored.`,
	"exit": `Shuts down the server.`,
	"user": `user save:
Saves important information about the user.

user password [password]:
Changes the stored password for the user.
`,
}
