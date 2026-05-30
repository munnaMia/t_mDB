package cmd

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/munnaMia/t_mDB/config"
	"github.com/munnaMia/t_mDB/internals/util"
	"github.com/munnaMia/t_mDB/parser"
	db "github.com/munnaMia/t_mDB/server/database"
)

func Run() {
	// load configurations
	cnf := config.GetConfig()

	// initialize inmemory database
	db := db.NewDatabase()

	// start tcp lintener
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(cnf.PORT))
	if err != nil {
		util.PrintError(string(parser.Encode(err)))
		os.Exit(1)
	}
	defer ln.Close()
	log.Println("Server running at PORT:", cnf.PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			util.PrintError(string(parser.Encode(err)))
			continue
		}

		go handleConnection(conn, db)
	}
}

func handleConnection(conn net.Conn, db *db.DB) {
	defer conn.Close()
	fmt.Println("test")
	var massage []byte

	reader := bufio.NewReader(conn)
	// writer := bufio.NewWriter(conn)

	// for {
	_, err := reader.Read(massage)
	if err != nil {
		util.PrintError(string(parser.Encode(err)))
		return
	}

	ackmsg := strings.ToUpper(strings.TrimSpace(string(massage)))
	res := fmt.Sprintf("Ack:  %s", ackmsg)
	if _, err := conn.Write([]byte(res)); err != nil {
		util.PrintError(string(parser.Encode(err)))
	}
	// }

}
