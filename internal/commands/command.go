package commands

type Command struct {
	Name        string
	Description string
	Run         func(args []string)
}

const (
	HELP string = "help"
	GET  string = "get"
	POST string = "post"
)

var AllCommands []Command

func init() {
	AllCommands = []Command{HelpCommand, GetCommand}
}
