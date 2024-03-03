package middleware

import (
	"reflect"
	"testing"

	"urlshortener/internal/app/config"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func TestLogger(t *testing.T) {
	type args struct {
		l   *zap.Logger
		cfg *config.Config
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Logger(tt.args.l, tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logger() = %v, want %v", got, tt.want)
			}
		})
	}
}
