package commands

import (
	"fmt"
)

var HelpCommand = Command{
	Name:        HELP,
	Description: "shows available commands",
	Run:         help,
}

func help(_ []string) {
	fmt.Println("zok - available commands:")
	fmt.Println("-------------------------")

	for _, command := range AllCommands {
		fmt.Printf("%s\t%s\n", command.Name, command.Description)
	}
}
