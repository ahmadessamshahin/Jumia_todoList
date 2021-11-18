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
	if ok := config(); ok {
		load([]Environment{&server{}, &Database{}})
	}
	fmt.Println("FINISH LOADING CONFIG 🚀🚀")
}

func config() bool {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(" 💀💀💀 fatal error config file: ", err)
		return false
	}
	return true
}

func load(envs []Environment) {
	for _, env := range envs {
		env.load()
		env.must()
		env.export()
	}
}
