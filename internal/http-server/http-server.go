package httpserver

import (
	"bytes"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"urlshortener/internal/config"
	"urlshortener/internal/storage"
)

type HTTPServer struct {
	cfg config.Config
	db  *storage.URLShortener
}

func New() (*HTTPServer, error) {
	// config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// db
	dbClient := storage.New()

	server := &HTTPServer{
		cfg: *cfg,
		db:  dbClient,
	}

	return server, nil
}

func (us *HTTPServer) Run() error {
	port := us.cfg.Port
	addr := fmt.Sprintf(":%s", port)

	mux := http.NewServeMux()
	mux.HandleFunc("/", us.MainPage)

	fmt.Println("Server is listening ...")
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		panic(err)
	}

	return nil
}

func (us *HTTPServer) MainPage(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/plain")

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
		us.db.Urls[id] = url

		// make response
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
		url, found := us.db.Urls[shortUrl]
		if !found {
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		// make response
		res.Header().Set("Location", url)
		res.WriteHeader(http.StatusTemporaryRedirect)
		return
	default:
		res.WriteHeader(http.StatusBadRequest)
		return
	}
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
