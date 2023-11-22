package handlers

import (
	"fmt"
	"strings"

	"github.com/izzyyhh/zok/utils/arrays"
)

func HandleArgs(args []string) {
	subcommand := strings.ToLower(args[0])

	if arrays.Contains(availableSubcommands, subcommand) {
		switch subcommand {
		case HELP:
			handleHelp()

		case GET:
			handleGet(args)
		}
	} else {
		fmt.Printf("subcommand '%s' is not supported", args[0])
	}
}
