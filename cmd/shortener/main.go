package main

import (
	"log"

	server "urlshortener/internal/app/http_server"
)

func main() {
	server, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
