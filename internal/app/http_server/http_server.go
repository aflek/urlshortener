package httpserver

import (
	"github.com/gin-gonic/gin"

	"urlshortener/internal/app/config"
	"urlshortener/internal/app/storage"
)

type UsServer struct {
	Cfg    config.Config
	DB     *storage.URLShortener
	Router *gin.Engine
}

func New() (*UsServer, error) {
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
	server := &UsServer{
		Cfg:    *cfg,
		DB:     dbClient,
		Router: router,
	}

	// preload routes with hendlers
	server.Routes()

	return server, nil
}

func (server *UsServer) Run() error {
	addr := server.Cfg.ServerAddress
	return server.Router.Run(addr)
}
