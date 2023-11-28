package handlers

import (
	"fmt"
	"strings"

	"github.com/izzyyhh/zok/internal/commands"
	"github.com/izzyyhh/zok/utils/arrays"
)

func HandleArgs(args []string) {
	subcommand := strings.ToLower(args[0])

	if arrays.Contains(availableSubcommands, subcommand) {
		switch subcommand {
		case HELP:
			commands.HelpCommand.Run(args)
		case GET:
			handleGet(args)

		case POST:
			handlePost(args)
		}

	} else {
		fmt.Printf("subcommand '%s' is not supported", args[0])
	}
}
