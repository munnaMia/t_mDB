package command

type CommandHandler func(args ...any) (error, []byte)

// all the commands
var CommandRegistry = map[string]CommandHandler{
	"PING": handlePing,

	"SET":    handleSet,
	"GET":    handleGet,
	"UP":     handleUp,
	"DEL":    handleDel,
	"EXISTS": handleExists,
}

// all the commands size
// (command key value) : 3
// (command key) : 2
var CommandSize = map[string]int{
	"PING": 1,

	"SET":    3,
	"GET":    2,
	"UP":     3,
	"DEL":    2,
	"EXISTS": 2,
}
