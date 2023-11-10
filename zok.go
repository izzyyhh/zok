package main

import (
	"fmt"
	"os"

	"github.com/izzyyhh/zok/internal/handlers"
)

func main() {
	args := os.Args[1:]
	argsCount := len(args)

	if argsCount == 0 {
		fmt.Println("zok is ready to make requests, for help use:\n'zok -h' or 'zok help'")
	} else if argsCount == 1 {
		handlers.HandleOneArg(args[0])
	} else {
		fmt.Println("TODO")
	}

}
