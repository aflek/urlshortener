package httpserver

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"urlshortener/internal/app/config"
	"urlshortener/internal/app/storage"
)

type UsApp struct {
	Cfg    config.Config
	Db     *storage.URLShortener
	Router *gin.Engine
}

func New() (*UsApp, error) {
	// config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// TODO logger

	// db
	dbClient := storage.New()

	// router
	router := gin.New()

	// init server data
	server := &UsApp{
		Cfg:    *cfg,
		Db:     dbClient,
		Router: router,
	}

	// preload routes with hendlers
	server.Routes()

	return server, nil
}

func (app *UsApp) Run() error {
	port := app.Cfg.Port
	addr := fmt.Sprintf(":%s", port)

	return app.Router.Run(addr)
}
