package tests

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"testing"

	server "urlshortener/internal/http-server"

	"github.com/stretchr/testify/assert"
)

const (
	PostURL = "http://localhost:8080"
)

func TestHendlers(t *testing.T) {
	var keyUrl string // key for short url
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
				contentType: "text/plain",
				statusCode:  http.StatusBadRequest,
			},
		},
		{
			name:    "POST request with long url",
			method:  http.MethodPost,
			request: "/",
			body:    "https://practicum.yandex.ru/",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusCreated,
			},
		},
		{
			name:    "GET request with short url",
			method:  http.MethodGet,
			request: "/",
			want: want{
				contentType: "text/plain",
				statusCode:  http.StatusTemporaryRedirect,
			},
		},
	}

	us, _ := server.New()

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// make body for request
			var reqBody io.Reader
			if test.body != "" {
				reqBody = strings.NewReader(test.body)
			} else {
				reqBody = nil
			}

			// define request
			if test.method == http.MethodGet && keyUrl != "" {
				test.request = fmt.Sprintf("/%s", keyUrl)
			}

			t.Log(test.request)

			req := httptest.NewRequest(test.method, test.request, reqBody)
			req.Header.Add("content-type", "text/plain")

			// response
			w := httptest.NewRecorder()

			// run handler
			us.MainPage(w, req)

			// handler result
			res := w.Result()

			// handler body
			resBody := w.Body.String()

			// get key for short url
			u, _ := url.Parse(resBody)
			keyUrl = path.Base(u.Path)

			assert.Equal(t, test.want.statusCode, res.StatusCode)
			assert.Equal(t, test.want.contentType, res.Header.Get("Content-Type"))
		})
	}
}
