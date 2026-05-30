package cmd

type CommandHandler func(args ...any) (error, []byte)

// all the commands
var CommandRegistry = map[string]CommandHandler{
	"PING":   handlePing,
	"SET":    handleSet,
	"GET":    handleGet,
	"UP":     handleUp,
	"DEL":    handleDel,
	"EXISTS": handleExists,
}
