package main

import (
	server "github.com/preegnees/simpauth/internal/controller/http"
)


func main() {

	if err := server.Run(); err != nil {
		panic(err)
	}
}
