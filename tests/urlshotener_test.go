package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	server "urlshortener/internal/app/http-server"
)

const (
	PostURL = "http://localhost:8080"
)

func TestHendlers(t *testing.T) {
	type want struct {
		contentType string
		statusCode  int
	}
	tests := []struct {
		name    string
		method  string
		body    string
		request string
		want    want
	}{
		{
			name:    "request with no body",
			method:  http.MethodPost,
			request: "/",
			body:    "",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "POST request with long url",
			method:  http.MethodPost,
			request: "/",
			body:    "https://practicum.yandex.ru/",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "GET request with short url",
			method:  http.MethodGet,
			request: "/",
			want: want{
				contentType: "text/plain; charset=utf-8",
				statusCode:  http.StatusTemporaryRedirect,
			},
		},
	}

	// init params
	var keyUrl string           // key for short url
	server, err := server.New() // server params
	if err != nil {
		t.Fatal(err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// make body for request
			reqBody := strings.NewReader(tt.body)

			// define request
			if tt.method == http.MethodGet && keyUrl != "" {
				tt.request = fmt.Sprintf("/%s", keyUrl)
			}

			// run handler
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.request, reqBody)
			server.Router.ServeHTTP(w, req)

			// handler result
			res := w.Result()
			resBody := w.Body.String()

			// get key for short url
			u, _ := url.Parse(resBody)
			keyUrl = path.Base(u.Path)

			// assert
			assert.Equal(t, tt.want.statusCode, w.Code)
			assert.Equal(t, tt.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
