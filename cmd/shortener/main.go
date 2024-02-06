package main

import (
	"flag"
	"log"
	"os"
	"strings"

	server "urlshortener/internal/app/http-server"
)

func main() {
	server, err := server.New()
	if err != nil {
		log.Fatal(err)
	}

	// Обработка данных из командной строки
	// пример: go run cmd/shortener/main.go -a=localhost:8888 -b=localhost:8080
	if len(os.Args) > 1 {
		// define flags
		a := flag.String("a", "", "server url (format: localhost:8888)")
		b := flag.String("b", "", "response url (format: localhost:8080)")

		flag.Parse()

		if a != nil {
			// get port
			s := strings.Split(*a, ":")
			if len(s) == 2 {
				server.Cfg.Port = s[1]
			}
		}

		if b != nil {
			server.Cfg.RedirectHost = *b
		}
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
