package main

import (
	"log"

	us "urlshortener/internal/http-server"
)

func main() {
	server, err := us.New()
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
