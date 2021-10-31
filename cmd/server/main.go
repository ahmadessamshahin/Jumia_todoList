package main

import (
	env "Jumia_todoList/config/environment"
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second)
		fmt.Println("running...", env.ServerEnv.Port)
	}
}
