package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "test BaseURL default value",
			want: "localhost:8080",
		},
		{
			name: "test ServerAddress default value",
			want: "localhost:8080",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := Load()
			assert.Equal(t, got.BaseURL, tt.want)
		})
	}
}
