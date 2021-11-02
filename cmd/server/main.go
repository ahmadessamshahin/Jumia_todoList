package main

import (
	"Jumia_todoList/api/handler"
	env "Jumia_todoList/config/environment"
	"log"
	"net/http"
)

func main() {
	handler.Handler.Serve()
	log.Fatal(http.ListenAndServe(env.ServerEnv.Port, handler.Handler.Router))
}
