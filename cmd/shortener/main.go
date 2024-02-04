package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type URLShortener struct {
	urls map[string]string
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 8

	rand.Seed(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}

func (us *URLShortener) MainPage(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:

		// get string with url from body of POST request text/plain
		buf := new(bytes.Buffer)

		body := req.Body
		_, err := buf.ReadFrom(body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		body.Close()
		url := strings.TrimSpace(buf.String())
		if url == "" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// id for short key
		id := generateShortKey()
		shortURL := fmt.Sprintf("http://%s/%s", req.Host, id)

		// save url
		us.urls[id] = url

		// make response
		res.Header().Set("content-type", "text/plain")
		res.WriteHeader(http.StatusCreated)
		res.Write([]byte(shortURL))
		return
	case http.MethodGet:
		// parse short url
		shortUrl := req.URL.Path[len("/"):]
		if shortUrl == "" {
			res.WriteHeader(http.StatusBadRequest)
			return
		}

		// restore url
		url, found := us.urls[shortUrl]
		if !found {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// make response
		res.Header().Set("content-type", "text/plain")
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
		return
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
}

func main() {
	port := "8080"
	addr := fmt.Sprintf(":%s", port)

	shortener := &URLShortener{
		urls: make(map[string]string),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", shortener.MainPage)

	fmt.Println("Server is listening ...")
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		panic(err)
	}
}
