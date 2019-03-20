package cli

import (
	"bufio"
	"fmt"
	"github.com/virepri/clipman-server/cli/commands"
	"github.com/virepri/clipman-server/shared"
	"os"
	"strings"
)

var reader *bufio.Reader

func StartCLI() {
	defer shared.ActiveServices.Done()

	fmt.Println("clipman-server version " + shared.Version)
	fmt.Println("Enter \"help\" for information on various commands.")

	reader = bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input := strings.Split(readLine(), " ")

		commands.Aliases[input[0]](input[1:])
	}
}

func readLine() string {
	if str, err := reader.ReadString('\n'); err == nil {
		return strings.Trim(str, "\n")
	} else {
		return ""
	}
}
