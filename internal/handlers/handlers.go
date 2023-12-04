package handlers

import (
	"fmt"
	"slices"
	"strings"

	"github.com/izzyyhh/zok/internal/commands"
)

func HandleArgs(args []string) {
	subcommand := strings.ToLower(args[0])
	idx := slices.IndexFunc(commands.AllCommands, func(c commands.Command) bool {
		return c.Name == subcommand
	})

	if idx != -1 {
		commands.AllCommands[idx].Run(args)
	} else {
		fmt.Println("subcommand not found...")
		commands.HelpCommand.Run(args)
	}

}
