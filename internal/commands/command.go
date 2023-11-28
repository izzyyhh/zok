package commands

import "fmt"

type Command struct {
	Name        string
	Description string
	Run         func(args []string)
}

var HelpCommand = Command{
	Name:        "help",
	Description: "help - shows available commands",
	Run:         help,
}

var allCommands []Command

func init() {
	allCommands = []Command{HelpCommand}
}

func help(_ []string) {
	for _, command := range allCommands {
		fmt.Println("zok - available commands:")
		fmt.Println(command.Description)
	}
}
