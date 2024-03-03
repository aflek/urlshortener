package httpserver

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type ShortenRq struct {
	URL string `json:"url"`
}
type ShortenRs struct {
	Result string `json:"result"`
}

// Create short url (json request/response)
func (server *UsServer) Shorten(c *gin.Context) {
	var (
		rq  ShortenRq
		rs  ShortenRs
		err error
	)

	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")

	defer func() {
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
	}()

	if err = c.ShouldBindJSON(&rq); err != nil {
		err = errors.New("request bind error")
		return
	}

	url := rq.URL
	if url == "" {
		err = errors.New("body is empty")
		return
	}

	// make id for short key
	id := generateShortKey()

	// make short url
	shortURL := fmt.Sprintf("%s/%s", server.Cfg.BaseURL, id)

	// save url (to map as tmp)
	server.DB.URLs[id] = url
	server.DB.UpdateFile(server.Cfg.FileStoragePath, id)

	// make response
	rs.Result = shortURL
	c.JSON(http.StatusCreated, rs)
}

// Create short url
func (server *UsServer) CreateShortURL(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")

	var err error
	defer func() {
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
	}()

	// get from request body string with 'long url'
	buf := new(bytes.Buffer)
	body := c.Request.Body
	_, err = buf.ReadFrom(body)
	if err != nil {
		return
	}
	body.Close()

	url := strings.TrimSpace(buf.String())
	if url == "" {
		err = errors.New("body is empty")
		return
	}

	// make id for short key
	id := generateShortKey()

	// make short url
	shortURL := fmt.Sprintf("%s/%s", server.Cfg.BaseURL, id)

	// save url (to map as tmp)
	server.DB.URLs[id] = url
	server.DB.UpdateFile(server.Cfg.FileStoragePath, id)

	c.String(http.StatusCreated, shortURL)
}

// Find short url
func (server *UsServer) FindShortURL(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/plain; charset=utf-8")

	var err error
	defer func() {
		if err != nil {
			c.Status(http.StatusBadRequest)
		}
	}()

	shortURL := c.Param("id")
	// restore url
	url, found := server.DB.URLs[shortURL]
	if !found {
		err = errors.New("url not found")
		return
	}

	c.Writer.Header().Set("Location", url)
	c.Status(http.StatusTemporaryRedirect)
}

func generateShortKey() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const keyLength = 8

	rand.NewSource(time.Now().UnixNano())
	shortKey := make([]byte, keyLength)
	for i := range shortKey {
		shortKey[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortKey)
}
