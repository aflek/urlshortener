package httpserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"urlshortener/internal/app/config"
	"urlshortener/internal/app/storage"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestShorten(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name    string
		bodyURL string
		want    want
	}{
		{
			name:    "POST /api/shorten with empty body",
			bodyURL: "",
			want: want{
				contentType: "application/json; charset=utf-8",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "POST /api/shorten with body",
			bodyURL: "https://practicum.yandex.ru",
			want: want{
				contentType: "application/json; charset=utf-8",
				statusCode:  http.StatusOK,
			},
		},
	}

	// config
	cfg, err := config.Load()
	if err != nil {
		return
	}

	// db
	dbClient := storage.New(cfg)

	// router
	router := gin.New()

	// init server
	server := &UsServer{
		Cfg:    cfg,
		DB:     dbClient,
		Router: router,
	}

	server.Router.POST("/api/shorten", server.Shorten)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// set request
			shortenRq := ShortenRq{
				URL: tt.bodyURL,
			}
			jsonValue, _ := json.Marshal(shortenRq)
			req, _ := http.NewRequest("POST", "/api/shorten", bytes.NewBuffer(jsonValue))

			// send request
			w := httptest.NewRecorder()
			server.Router.ServeHTTP(w, req)

			// result
			result := w.Result()
			defer result.Body.Close()

			// parse response body
			var resp ShortenRs
			json.Unmarshal(w.Body.Bytes(), &resp)

			// test results
			assert.Equal(t, tt.want.statusCode, w.Code)
			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))
			if w.Code == http.StatusOK {
				assert.NotEmpty(t, resp.Result)
			}
		})
	}
}
