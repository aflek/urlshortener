package httpserver

import (
	//"compress/gzip"
	"github.com/gin-contrib/gzip"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"urlshortener/internal/app/config"
	"urlshortener/internal/app/logger"
	"urlshortener/internal/app/middleware"
	"urlshortener/internal/app/storage"
)

type UsServer struct {
	Cfg    *config.Config
	Log    *zap.Logger
	DB     *storage.URLShortener
	Router *gin.Engine
}

func New() (*UsServer, error) {
	// config
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	// logger
	logger, err := logger.New(cfg)
	if err != nil {
		return nil, err
	}

	// db
	dbClient := storage.New()

	// router
	router := gin.New()

	// middleware
	router.Use(middleware.Logger(logger, cfg))
	// не могу разобраться почему моя функция не перезаписывает заголовок gin
	// временно ее выключил:
	// router.Use(middleware.Gzip(gzip.BestSpeed))
	// а вот из пакета, работает и аналогичным образом перезаписывает:
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	// init server data
	server := &UsServer{
		Cfg:    cfg,
		Log:    logger,
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
