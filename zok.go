package main

import (
	"fmt"
	"os"

	"github.com/izzyyhh/zok/internal/handlers"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("zok is ready to make requests, for help use:\n'zok help'")
	} else {
		handlers.HandleArgs(args)
	}

}
