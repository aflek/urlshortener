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
			want: "http://localhost:8080",
		},
		{
			name: "test ServerAddress default value",
			want: "localhost:8080",
		},
	}

	got, _ := Load()

	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			switch i {
			case 0:
				assert.Equal(t, got.BaseURL, tt.want)
			case 1:
				assert.Equal(t, got.ServerAddress, tt.want)
			}
		})
	}
}
