package rest

import (
	"fmt"
)

type Server struct {
	IP   string
	Port int
	Name string
}

func (c *Server) GetAddress() string {
	return fmt.Sprintf("http://%s:%d", c.IP, c.Port)
}
