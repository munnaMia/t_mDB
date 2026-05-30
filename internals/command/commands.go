package command

import db "github.com/munnaMia/t_mDB/server/database"

type CommandHandler func(db *db.DB, args ...any) (error, []byte)

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
