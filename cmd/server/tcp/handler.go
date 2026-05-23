package tcp

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println(conn)
}
