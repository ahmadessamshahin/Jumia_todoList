package env

import (
	"fmt"
	"github.com/spf13/viper"
)

type Environment interface {
	must() error
	load()
	export()
}

func init() {
	fmt.Println("START LOADING CONFIG ....")
	config()
	load([]Environment{&server{}, &database{}})
	fmt.Println("FINISH LOADING CONFIG 🚀🚀")
}

func config() {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}

func load(envs []Environment) {
	for _, env := range envs {
		env.load()
		env.must()
		env.export()
	}
}
