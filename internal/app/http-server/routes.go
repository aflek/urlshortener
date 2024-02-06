package httpserver

func (server *UsApp) Routes() {
	server.Router.POST("/", server.CreateShortURL)
	server.Router.GET("/:id", server.FindShortURL)
}
