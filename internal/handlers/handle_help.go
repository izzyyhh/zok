package handlers

import "fmt"

const instructions = `Usage: zok [http method] [flags] <url>
zok supports following http methods:
	- GET
	- POST

zok has following flags:
	POST:   --body`

func handleHelp() {
	fmt.Println(instructions)
}
