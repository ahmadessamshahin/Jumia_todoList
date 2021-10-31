package env

import (
	"fmt"
	"os"
)

var ServerEnv server

type server struct {
	Port string
}

func (s server) must() error {
	if s.Port == "" {
		panic(fmt.Errorf("MISSING SERVER INITIALIZATION PORT"))
	}
	return nil
}

func (s *server) load() {
	s.Port = os.Getenv("APP_PORT")
}

func (s server) export() {
	ServerEnv = s
}
