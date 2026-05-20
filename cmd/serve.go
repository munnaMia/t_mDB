package cmd

import (
	"log"
	"net"
)

func Run() {
	svr := NewServer(6666)

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
