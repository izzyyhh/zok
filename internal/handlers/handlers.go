package handlers

import (
	"fmt"

	"github.com/izzyyhh/zok/internal/httpclient"
	"github.com/izzyyhh/zok/utils/arrays"
	"github.com/izzyyhh/zok/utils/validate"
)

var availableSubcommands = []string{"help", "test"}

//check if is subcommand
//if no execute http get
//otherwise handlesubcommand

func HandleOneArg(arg string) {

	if arrays.Contains(availableSubcommands, arg) {
		if arg == "help" {
			handleHelp()
			return
		}

		// can i do a dictionary with pointers to methods? {help: handleHelp, test: handleTtest}
	}

	validation := validate.Url(arg)

	if validation {
		result := httpclient.Get(arg)
		fmt.Println(result)
	} else {
		fmt.Println("invalid url")
	}
}

func handleHelp() {
	fmt.Println("HELPING YOU :)")
}
