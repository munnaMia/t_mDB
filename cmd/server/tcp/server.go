package tcp

import "fmt"

type Server struct {
	PORT int
}

func NewServer(port int) *Server {
	return &Server{
		PORT: port,
	}
}

func (svr *Server) addr() string {
	return fmt.Sprintf(":%d", svr.PORT)
}
