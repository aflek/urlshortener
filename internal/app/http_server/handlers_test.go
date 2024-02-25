package httpserver

import (
	"testing"

	"github.com/gin-gonic/gin"

	"urlshortener/internal/app/config"
	"urlshortener/internal/app/storage"
)

func TestUsServer_CreateShortURL(t *testing.T) {
	type fields struct {
		Cfg    config.Config
		DB     *storage.URLShortener
		Router *gin.Engine
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// test cases after adding DB.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &UsServer{
				Cfg:    tt.fields.Cfg,
				DB:     tt.fields.DB,
				Router: tt.fields.Router,
			}
			server.CreateShortURL(tt.args.c)
		})
	}
}
