package commands

type Command struct {
	Cmd byte
	Args []string
}

var Aliases = map[byte]func(Args []string){}

/*
Command structure:
[CMD, 10, ...., 10, ...., 0]
Every ASCII LF (\n) character is a new argument.
Every command ends with a NULL byte.
 */
func ParseCmd(buffer []byte) Command {
	cmd := Command{ Cmd:buffer[0], Args:make([]string, 0) }

	start := -1
	k := -1
	var v byte
	for k,v = range buffer {
		if k != 0 && v == 0 {
			if start != -1 {
				cmd.Args = append(cmd.Args, string(buffer[start:k]))
			}
			break
		}

		if k != 0 && v == 10 { //LF ASCII
			if start != -1 {
				cmd.Args = append(cmd.Args, string(buffer[start:k]))
			}
			start = k+1
		}
	}

	return cmd
}
