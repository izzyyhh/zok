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

	urlValidation := validate.Url(arg)

	if urlValidation {
		body, err := httpclient.Get(arg)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println(string(body))

	} else {
		fmt.Println("invalid url")
	}
}

func handleHelp() {
	const instructions = `
Usage: zok [http method] [flags] <url>
zok supports following http methods: 
	- GET (default)
	- POST

zok has following flags:
	POST:	-d, --data`

	fmt.Println(instructions)
}

// TODO add flags with flag package, fisrt one data for post req
