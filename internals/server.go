package internals

import "fmt"

type Server struct {
	PORT int
}

func NewServer(port int) *Server {
	return &Server{
		PORT: port,
	}
}

// return PORT with : colon. Like :6666
func (svr *Server) GetPORT() string {
	return fmt.Sprintf(":%d", svr.PORT)
}
