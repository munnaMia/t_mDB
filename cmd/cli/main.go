package main

import (
	"fmt"
	"net"

	"github.com/munnaMia/t_mDB/config"
)

func main() {
	// load configurations
	cnf := config.GetConfig()

	address := fmt.Sprintf("%s:%d", cnf.IP, cnf.PORT)

	net.Dial("tcp", address)
}
