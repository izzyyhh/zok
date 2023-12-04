package commands

import "fmt"

var HelpCommand = Command{
	Name:        "help",
	Description: "help - shows available commands",
	Run:         help,
}

func help(_ []string) {
	for _, command := range AllCommands {
		fmt.Println("zok - available commands:")
		fmt.Println(command.Description)
	}
}
