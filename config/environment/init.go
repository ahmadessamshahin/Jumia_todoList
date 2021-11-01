package env

import (
	"fmt"
)

type Environment interface {
	must() error
	load()
	export()
}

func init() {
	fmt.Println("START LOADING CONFIG ....")
	load([]Environment{&server{}, &database{}})
	fmt.Println("FINISH LOADING CONFIG 🚀🚀")
}

func load(envs []Environment) {
	for _, env := range envs {
		env.load()
		_ = env.must()
		env.export()
	}
}
