package cmd

import (
	"log"
	"net"

	"github.com/munnaMia/t_mDB/config"
	"github.com/munnaMia/t_mDB/internals"
)

func Run() {
	// load configurations
	cnf := config.GetConfig()

	// init new server struct
	svr := internals.NewServer(cnf.PORT)

	// start tcp lintener
	ln, err := net.Listen("tcp", svr.GetPORT())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server running at PORT:", svr.PORT)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handleConnection(conn)
	}
}
