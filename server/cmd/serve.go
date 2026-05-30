package cmd

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/munnaMia/t_mDB/config"
	"github.com/munnaMia/t_mDB/internals/util"
	"github.com/munnaMia/t_mDB/parser"
)

func Run() {
	// load configurations
	cnf := config.GetConfig()

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

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("test")
	// var massage []byte

	// reader := bufio.NewReader(conn)
	// // writer := bufio.NewWriter(conn)

	// for {
	// 	_, err := reader.Read(massage)
	// 	if err != nil {

	// 		return
	// 	}
	// }

}
