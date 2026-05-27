package cmd

type CommandHandler func(args ...any) (error, []byte)

// all the commands
var CommandRegistry = map[string]CommandHandler{
	"SET": handleSet,
}
