package tcp

import (
	"log"
	"net"

	"github.com/munnaMia/t_mDB/config"
)

func Run() {
	// load configurations
	cnf := config.GetConfig()

	// init new server struct
	svr := NewServer(cnf.PORT)

	// start tcp lintener
	ln, err := net.Listen("tcp", svr.addr())
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
