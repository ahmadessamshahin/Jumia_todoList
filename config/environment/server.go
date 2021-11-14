package env

import (
	"fmt"

	"github.com/spf13/viper"
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
	s.Port = viper.Get("APP_PORT").(string)
}

func (s server) export() {
	ServerEnv = s
}
