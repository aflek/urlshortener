package logger

import (
	"reflect"
	"testing"

	"urlshortener/internal/app/config"

	"go.uber.org/zap"
)

func TestNew(t *testing.T) {
	type args struct {
		srvConf *config.Config
	}
	tests := []struct {
		name    string
		args    args
		want    *zap.Logger
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.srvConf)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
